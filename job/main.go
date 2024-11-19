package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	v1 "genie-ai-skywalking/api/v1"
	"genie-ai-skywalking/internal/common"
	"genie-ai-skywalking/internal/conf"
	"genie-ai-skywalking/internal/entity"
	"genie-ai-skywalking/internal/pkg/dayu_trace"
	"genie-ai-skywalking/pkg/utils"
	"genie-ai-skywalking/pkg/zlog"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	nacosConfig "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-resty/resty/v2"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/satori/go.uuid"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type QueryListReq struct {
	AppId    string `json:"app_id,omitempty"`
	DeviceId string `json:"device_id,omitempty"`
	Origin   string `json:"origin,omitempty"`
	Function string `json:"function,omitempty"`
	TraceId  string `json:"trace_id,omitempty"`
	FromTime string `json:"from_time,omitempty"`
	ToTime   string `json:"to_time,omitempty"`
	Page     int64  `json:"page,omitempty"`
	PageSize int64  `json:"page_size,omitempty"`
}
type QueryDetailReq struct {
	TraceId  string `json:"trace_id,omitempty"`
	FromTime string `json:"from_time,omitempty"`
	ToTime   string `json:"to_time,omitempty"`
	Page     int64  `json:"page,omitempty"`
	PageSize int64  `json:"page_size,omitempty"`
}

type SlsClient struct {
	Client   *sls.Client
	Project  string
	Logstore string
}

func NewSlsClient(conf *conf.Data) *SlsClient {
	return &SlsClient{
		Client: &sls.Client{
			Endpoint:        conf.Sls.Endpoint,
			AccessKeyID:     conf.Sls.AccessKeyId,
			AccessKeySecret: conf.Sls.AccessKeySecret,
		},
		Project:  conf.Sls.Project,
		Logstore: conf.Sls.Logstore,
	}
}

var (
	// Name is the name of the compiled software.
	Name = "genie-sky-walking" // TODO:默认为当前项目名称，会影响服务发现注册等功能，一旦项目发布不要后续修改，除非知道自己在做什么！！！
	// Version is the version of the compiled software.
	Version = "1.0.0"
	// App Env
	env string
	// flagconf is the config flag.
	flagconf         string
	nacos            bool
	nacosLogDir      string
	nacosCacheDir    string
	nacosServer      string
	nacosPort        uint64
	nacosNamespaceId string
	nacosGroupId     = "genie-ai-skywalking" //TODO:edit me
	dataIds          = []string{             //TODO:edit me while add configs
		"config.yaml",
	}
	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&env, "env", string(common.LocalEnv), "app env, eg: -env local")
	flag.StringVar(&Name, "name", Name, "app name, eg: -name kratos-layout")
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf=config.yaml")
	flag.BoolVar(&nacos, "nacos", false, "use nacos, eg: -nacos=true")
	flag.StringVar(&nacosServer, "nacosServer", "127.0.0.1", "nacos host, eg: -nacosServer=127.0.0.1")
	flag.Uint64Var(&nacosPort, "nacosPort", 8848, "nacos port, eg: -nacosPort 8488")
	flag.StringVar(&nacosNamespaceId, "nacosNamespaceId", "namespaceId", "nacos namespaceId, eg: -nacosNamespaceId=id")
	flag.StringVar(&nacosCacheDir, "nacosCacheDir", "./logs", "nacos cacheDir, eg: -nacosCacheDir=./logs")
	flag.StringVar(&nacosLogDir, "nacosLogDir", "./logs", "nacos logDir, eg: -nacoslogDir=./logs")
}
func NewConfigClient(address string, port uint64, namespaceId string, logDir string, cacheDir string) config_client.IConfigClient {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(address, port),
	}
	cc := &constant.ClientConfig{
		NamespaceId:         namespaceId, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              logDir,
		CacheDir:            cacheDir,
		LogLevel:            "debug",
	}
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	return client
}

func main() {
	flag.Parse()
	var c config.Config
	if nacos {
		nacosClient := NewConfigClient(nacosServer, nacosPort, nacosNamespaceId, nacosLogDir, nacosCacheDir)
		configSources := make([]config.Source, 0)
		for _, dataId := range dataIds {
			configSources = append(configSources, nacosConfig.NewConfigSource(nacosClient, nacosConfig.WithGroup(nacosGroupId), nacosConfig.WithDataID(dataId)))
		}
		c = config.New(
			config.WithSource(configSources...),
		)
	} else {
		c = config.New(
			config.WithSource(
				file.NewSource(flagconf),
			),
		)
	}
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	InitQueryTraceTasks1(&bc)

	//	InitQueryTraceTasks(&bc)
	//args := os.Args
	//fmt.Println(args)
	//if len(args) > 2 {
	//	strs := strings.Split(args[2], "=")
	//	if len(strs) == 2 {
	//		switch strs[1] {
	//		case "zuowen_trace":
	//			InitQueryTraceTasks(&bc)
	//		}
	//	}
	//}
}

type QueryEssayTasks struct {
	TraceId  string `json:"trace_id"`
	DeviceId string `json:"device_id"`
	Logs     string `json:"logs"`
}
type Task struct {
	Id     primitive.ObjectID     `bson:"_id,omitempty"`
	MainId primitive.ObjectID     `bson:"main_id"`
	Type   int                    `bson:"type"`
	Info   map[string]interface{} `bson:"info"`
}

func InitEssayDataServiceDB() (*gorm.DB, error) {
	dsn := "znyj_bi_ro:7^zjTPJuAyEMfnzs$RtTnJCr^F8*JBrC@tcp(znyj-bi.t596s9xa7ihog.oceanbase.aliyuncs.com:3306)/znyj_bi?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
}
func InitDataServiceDB() (*gorm.DB, error) {
	dsn := "znyj_bi_ro:7^zjTPJuAyEMfnzs$RtTnJCr^F8*JBrC@tcp(znyj-bi.t596s9xa7ihog.oceanbase.aliyuncs.com:3306)/znyj_bi?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
}
func InitQueryTraceTasks2(bc *conf.Bootstrap) {
	essayDB, err := InitEssayDataServiceDB()
	if err != nil {
		fmt.Println("db链接出错, error: ", err.Error())
		return
	}

	//page := 0
	//for {
	//	essayDB, err := InitEssayDataServiceDB()
	//	if err != nil {
	//		fmt.Println("db链接出错, error: ", err.Error())
	//		return
	//	}
	//	page = page + 1
	//	var essayResp []QueryEssayTasks
	//	err = essayDB.Table("query_sls_trace").Select("trace_id, device_id, logs").Limit(1000).Offset((page - 1) * 1000).Find(&essayResp).Error
	//	if err != nil {
	//		fmt.Println("db查询出错, error: ", err.Error())
	//		return
	//	}
	//	if len(essayResp) == 0 {
	//		break
	//	}
	//	var wg sync.WaitGroup
	//	wg.Add(len(essayResp))
	//	for _, logMap := range essayResp {
	//		go func(logMap QueryEssayTasks) {
	//			defer wg.Done()
	//			var logs []entity.Log
	//			_ = json.Unmarshal([]byte(logMap.Logs), &logs)
	//			var input string
	//			for _, l := range logs {
	//				if l.Name == "input" {
	//					input = l.LogAttribute.Key
	//				}
	//			}
	//			//ToDo 调用三方接口
	//			err = UpdateButterfly(context.Background(), logMap.TraceId, logMap.DeviceId, input, bc)
	//			if err != nil {
	//				return
	//			}
	//		}(logMap)
	//
	//	}
	//	wg.Wait()
	//}
	var essayResp []QueryEssayTasks
	err = essayDB.Table("query_sls_trace").Select("trace_id, device_id, logs").Limit(1).Find(&essayResp).Error
	if err != nil {
		fmt.Println("db查询出错, error: ", err.Error())
		return
	}
	var logger log.Logger
	var c = NewClient(bc.Services.Butterfly, logger)

	for _, logMap := range essayResp {
		var logs []entity.Log
		_ = json.Unmarshal([]byte(logMap.Logs), &logs)
		var input string
		for _, l := range logs {
			if l.Name == "input" {
				input = l.LogAttribute.Key
			}
		}
		//ToDo 调用三方接口
		err = UpdateButterfly(context.Background(), logMap.TraceId, logMap.DeviceId, input, bc, c)
		if err != nil {
			return
		}
	}

}

func InitQueryTraceTasks1(bc *conf.Bootstrap) {
	//协程处理
	//ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	//defer cancel()

	//clientOptions := options.Client().ApplyURI("mongodb://root:pbH7w58ll1UQZb1_@s-2ze411bc75fcfab4.mongodb.rds.aliyuncs.com:3717/")
	//
	//// 连接到MongoDB
	//client, err := mongo.Connect(context.TODO(), clientOptions)
	//
	//if err != nil {
	//	fmt.Println("mdb链接出错, error: ", err.Error())
	//	return
	//}
	//
	//// 检查连接
	//err = client.Ping(context.TODO(), nil)
	//
	//if err != nil {
	//	fmt.Println("mdb链接出错, error: ", err.Error())
	//	return
	//}
	//
	//fmt.Println("Connected to MongoDB!")
	//
	//// 在此处进行其他操作，例如插入文档、查询文档等
	////opts := options.Find()
	//collection := client.Database("xiaosi_butterfly").Collection("main_tasks")
	//cond := bson.M{}
	//cond["type"] = 14
	//
	//cursor, err := collection.Find(context.TODO(), cond)
	//if err != nil {
	//
	//}
	//var tasks []Task
	//err = cursor.All(context.Background(), &tasks)
	//var tasksTraceId []string
	//for _, v := range tasks {
	//	s := v.Info["trace_id"]
	//	fmt.Println(s)
	//	tasksTraceId = append(tasksTraceId, cast.ToString(s))
	//}
	//// 断开连接
	//err = client.Disconnect(context.TODO())
	//
	//if err != nil {
	//	fmt.Println("mdb链接断开出错, error: ", err.Error())
	//	return
	//}

	essayDB, err := InitEssayDataServiceDB()
	if err != nil {
		fmt.Println("db链接出错, error: ", err.Error())
		return
	}

	//page := 0
	//for {
	//	essayDB, err := InitEssayDataServiceDB()
	//	if err != nil {
	//		fmt.Println("db链接出错, error: ", err.Error())
	//		return
	//	}
	//	page = page + 1
	//	var essayResp []QueryEssayTasks
	//	err = essayDB.Table("query_sls_trace").Select("trace_id, device_id, logs").Limit(1000).Offset((page - 1) * 1000).Find(&essayResp).Error
	//	if err != nil {
	//		fmt.Println("db查询出错, error: ", err.Error())
	//		return
	//	}
	//	if len(essayResp) == 0 {
	//		break
	//	}
	//	var wg sync.WaitGroup
	//	wg.Add(len(essayResp))
	//	for _, logMap := range essayResp {
	//		go func(logMap QueryEssayTasks) {
	//			defer wg.Done()
	//			var logs []entity.Log
	//			_ = json.Unmarshal([]byte(logMap.Logs), &logs)
	//			var input string
	//			for _, l := range logs {
	//				if l.Name == "input" {
	//					input = l.LogAttribute.Key
	//				}
	//			}
	//			//ToDo 调用三方接口
	//			err = UpdateButterfly(context.Background(), logMap.TraceId, logMap.DeviceId, input, bc)
	//			if err != nil {
	//				return
	//			}
	//		}(logMap)
	//
	//	}
	//	wg.Wait()
	//}
	var essayResp []QueryEssayTasks
	err = essayDB.Table("query_sls_trace").Select("trace_id, device_id, logs").Find(&essayResp).Error
	if err != nil {
		fmt.Println("db查询出错, error: ", err.Error())
		return
	}
	//if len(essayResp) == 0 {
	//	break
	//}
	var logger log.Logger
	var c = NewClient(bc.Services.Butterfly, logger)

	idx := 0
	mtx := &sync.Mutex{}
	failedIdList := make([]int, 0, 10000)
	var worker = func(id int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			mtx.Lock()
			if idx >= len(essayResp) {
				mtx.Unlock()
				break
			}
			curIndex := idx
			idx++
			mtx.Unlock()
			var logs []entity.Log
			_ = json.Unmarshal([]byte(essayResp[curIndex].Logs), &logs)

			var input string
			for _, l := range logs {
				if l.Name == "input" {
					input = l.LogAttribute.Key
				}
			}
			err = UpdateButterfly(context.Background(), essayResp[curIndex].TraceId, essayResp[curIndex].DeviceId, input, bc, c)
			if err != nil {
				failedIdList = append(failedIdList, curIndex)
				continue
			}
		}
	}
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go worker(i, wg)
	}
	wg.Wait()
	failedIdFile, err := os.Create("./failed_id.txt")
	if err != nil {
		fmt.Println("Error creating")
		return
	}
	defer failedIdFile.Close()
	writer := bufio.NewWriter(failedIdFile)
	for _, failedId := range failedIdList {
		var logs []entity.Log
		_ = json.Unmarshal([]byte(essayResp[failedId].Logs), &logs)

		var input string
		for _, l := range logs {
			if l.Name == "input" {
				input = l.LogAttribute.Key
			}
		}
		err = UpdateButterfly(context.Background(), essayResp[failedId].TraceId, essayResp[failedId].DeviceId, input, bc, c)
		if err != nil {
			writer.WriteString(fmt.Sprintf("%d\n", failedId))
			continue
		}
	}
	writer.Flush()
}

//func InitQueryTraceTasks(bc *conf.Bootstrap) {
//	//协程处理
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	query := QueryListReq{
//		Origin:   "TextCorrection",
//		Function: "proxy_text-correction",
//		FromTime: "2023-12-11 00:00:00",
//		Page:     1,
//		PageSize: 40,
//		ToTime:   "2023-12-17 00:00:00",
//	}
//
//	list, err := QueryList(ctx, &query, bc)
//	if err != nil {
//		return
//	}
//	if list.Total == 0 {
//		fmt.Println("没有找到任何数据")
//	}
//	//获取数据"2023-11-11 17:03:35",
//	//			ToTime:   "2023-12-11 17:03:35",
//	for _, v := range list.List {
//		q := QueryDetailReq{
//			TraceId:  v.TraceId,
//			FromTime: "2023-12-01 00:00:00",
//			ToTime:   "2023-12-10 00:00:00",
//			Page:     1,
//			PageSize: 10000,
//		}
//		detail, err := QueryDetail(ctx, &q, bc)
//		if err != nil {
//			return
//		}
//		if len(detail.List) == 0 {
//			continue
//		}
//		//ToDo 调用三方接口
//		err = UpdateButterfly(ctx, detail.List[0].TraceId, detail.List[0].DeviceId, detail.List[0].Input, bc)
//		if err != nil {
//			return
//		}
//	}
//	//调用三方接口
//
//	//判断返回
//}

type Client struct {
	log    *log.Helper
	client *resty.Client
}

type CommonRespField struct {
	ErrReason  string `json:"error_reason"`
	ErrMsg     string `json:"error_msg"`
	ServerTime int    `json:"server_time"`
	TraceId    string `json:"trace_id"`
}

type UpdateQuestionResp struct {
	CommonRespField
	Data TaskCheck `json:"data"`
}

type TaskCheck struct {
	TaskId string `json:"task_id"`
}

func getHeaders(c *conf.ServiceBaseConfig) map[string]string {
	timestamp := cast.ToString(time.Now().Unix())
	nonce := uuid.NewV4().String()
	md5Str := fmt.Sprintf("%s&X-Genie-Timestamp=%s&X-Genie-Nonce=%s&X-Genie-DeviceId=%s", c.Secret, timestamp, nonce, "1234wadwadd567")
	var sign = utils.MD5(md5Str)
	return map[string]string{
		"X-Genie-DeviceId":  "1234wadwadd567",
		"X-Genie-Timestamp": timestamp,
		"X-Genie-Nonce":     nonce,
		"X-Genie-AppId":     c.Key,
		"X-Genie-Sign":      sign,
		"X-Genie-Version":   "1.0.0",
		"X-Genie-Platform":  "iOS",
		"X-Genie-OsVersion": "15.7",
	}
}
func NewClient(config *conf.ServiceBaseConfig, logger log.Logger) *Client {
	//timestamp := cast.ToString(time.Now().Unix())
	//nonce := uuid.NewV4().String()
	//md5Str := fmt.Sprintf("%s&X-Genie-Timestamp=%s&X-Genie-Nonce=%s&X-Genie-DeviceId=%s", config.Secret, timestamp, nonce, "1234wadwadd567")
	//var sign = utils.MD5(md5Str)
	client := resty.New()
	client.SetBaseURL(config.Host).
		SetRetryCount(2).
		SetHeader("Content-Type", "application/json").
		SetTimeout(config.Timeout.AsDuration()).
		SetRetryWaitTime(time.Millisecond * 100).
		OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
			// 从ctx获取request-id
			requestId := trace.SpanContextFromContext(r.Context()).TraceID().String()
			r.SetHeaders(map[string]string{
				//"X-Genie-DeviceId":  "1234wadwadd567",
				//"X-Genie-Timestamp": timestamp,
				//"X-Genie-Nonce":     nonce,
				//"X-Genie-AppId":     config.Key,
				//"X-Genie-Sign":      sign,
				//"X-Genie-Version":   "1.0.0",
				//"X-Genie-Platform":  "iOS",
				//"X-Genie-OsVersion": "15.7",
				//"Authorization":     "Bearer",
				"request-id": requestId,
			})
			return nil
		})
	return &Client{
		log:    log.NewHelper(logger),
		client: client,
	}
}

type Picture struct {
	Pics       []string `json:"answer_url"`
	GradeId    int      `json:"grade"`
	LargeModel int      `json:"large_model"`
}

//var c = NewClient(bc.Services.Butterfly, logger)

func UpdateButterfly(ctx context.Context, traceId, Sn, input string, bc *conf.Bootstrap, c *Client) (err error) {

	var pics Picture
	//fmt.Printf(input)
	input = "\"" + input + "\""
	//fmt.Printf(input)
	s, err := strconv.Unquote(input)
	if err != nil {
		return err
	}
	if input != "" {
		err := json.Unmarshal([]byte(s), &pics)
		if err != nil {
			return err
		}
	}
	req := map[string]interface{}{
		"trace_id": traceId,
		"sn":       Sn,
		"picture":  pics.Pics,
		"grade_id": pics.GradeId,
	}
	result := &UpdateQuestionResp{}
	res, err := c.client.SetRetryCount(3).SetTimeout(15 * time.Second).R().
		SetResult(result).
		SetBody(req).
		SetContext(ctx).
		SetError(&result).
		Post("/inner/tasks/add-essay")
	if err != nil || !res.IsSuccess() || result.ErrReason != "success" {
		c.log.WithContext(ctx).Warnf("reject task , res:%+v, result:%+v", res, result)
		err = errors.New(result.ErrMsg)
		return
	}
	fmt.Println("success======", traceId)
	return
}

func convertConfigToMap(conf *conf.Enums) (appMap map[string]string, originMap map[string]string, functionMap map[string]string) {

	appMap = make(map[string]string, 0)
	originMap = make(map[string]string, 0)
	functionMap = make(map[string]string, 0)

	for _, v := range conf.App {
		appMap[v.Id] = v.Name
	}
	for _, v := range conf.Origin {
		originMap[v.Id] = v.Name
	}
	for _, v := range conf.Function {
		functionMap[v.Id] = v.Name
	}
	return
}

func GetOrDefault(m map[string]string, k string) string {
	if v, ok := m[k]; ok {
		return v
	}
	return k
}

func QueryList(ctx context.Context, req *QueryListReq, bc *conf.Bootstrap) (*v1.QueryListReply, error) {

	//var config *conf.Enums
	appMap, originMap, functionMap := convertConfigToMap(bc.Enums)
	lines := req.PageSize
	if lines <= 0 || lines > 100 {
		lines = 20
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	offset := (req.Page - 1) * lines
	from, _ := utils.ParseDateTimeInSec(req.FromTime)
	to, _ := utils.ParseDateTimeInSec(req.ToTime)

	var queryBuilder strings.Builder
	_, _ = fmt.Fprint(&queryBuilder, "* and service: paas-proxy")
	if len(req.AppId) > 0 {
		_, _ = fmt.Fprint(&queryBuilder, " and attribute.appId: ", req.AppId)
	}
	if len(req.DeviceId) > 0 {
		_, _ = fmt.Fprint(&queryBuilder, " and attribute.deviceId: ", req.DeviceId)
	}
	if len(req.Origin) > 0 {
		_, _ = fmt.Fprint(&queryBuilder, " and attribute.origin: ", req.Origin)
	}
	if len(req.Function) > 0 {
		_, _ = fmt.Fprint(&queryBuilder, " and attribute.function: ", req.Function)
	}
	if len(req.TraceId) > 0 {
		_, _ = fmt.Fprint(&queryBuilder, " and traceID: ", req.TraceId)
	}

	logReq := &sls.GetLogRequest{
		From:    from,
		To:      to,
		Lines:   lines,
		Offset:  offset,
		Reverse: true,
		Query:   queryBuilder.String(),
	}

	totalCount := int64(0)
	//var con *conf.Data
	slsC := NewSlsClient(bc.Data)
	for {
		// GetHistograms API Ref: https://intl.aliyun.com/help/doc-detail/29030.htm
		ghResp, err := slsC.Client.GetHistograms(slsC.Project, slsC.Logstore, "", logReq.From, logReq.To, logReq.Query)
		if err != nil {
			fmt.Printf("GetHistograms fail, err: %v\n", err)
			time.Sleep(10 * time.Millisecond)
			continue
		}
		//fmt.Printf("complete: %s, count: %d, histograms: %v\n", ghResp.Progress, ghResp.Count, ghResp.Histograms)
		totalCount += ghResp.Count
		if ghResp.Progress == "Complete" {
			break
		}
	}

	logResp, err := slsC.Client.GetLogsV2(slsC.Project, slsC.Logstore, logReq)
	if err != nil {
		return nil, err
	}
	//fmt.Println(logResp)

	replyList := make([]*v1.QueryListReplyItem, 0)
	for _, logMap := range logResp.Logs {
		attributeStr := logMap["attribute"]
		//fmt.Println("attributeStr: ", attributeStr)
		var attribute entity.Attribute
		_ = json.Unmarshal([]byte(attributeStr), &attribute)
		start, _ := strconv.ParseInt(logMap["start"], 10, 64)
		startTime := utils.FormatDateTime(time.UnixMicro(start))
		end, _ := strconv.ParseInt(logMap["end"], 10, 64)
		endTime := utils.FormatDateTime(time.UnixMicro(end))
		duration := end - start
		replyItem := &v1.QueryListReplyItem{
			AppId:           GetOrDefault(appMap, attribute.AppId),
			DeviceId:        attribute.DeviceId,
			Function:        GetOrDefault(functionMap, attribute.Function),
			FunctionVersion: attribute.FunctionVersion,
			Origin:          GetOrDefault(originMap, attribute.Origin),
			TraceId:         logMap["traceID"],
			StartTime:       startTime,
			EndTime:         endTime,
			Duration:        duration,
		}
		replyList = append(replyList, replyItem)
	}
	return &v1.QueryListReply{
		List:  replyList,
		Total: totalCount,
	}, nil
}

func initConfig() (confData *conf.Data, logger log.Logger) {
	c := config.New(
		config.WithSource(
			file.NewSource("../../configs"),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	confData = bc.Data
	zlog.Init("data_test", bc.Log.Filename, int(bc.Log.MaxSize), int(bc.Log.MaxBackup), int(bc.Log.MaxAge), bc.Log.Compress)
	defer func() {
		_ = zlog.Sync()
	}()
	logger = log.With(zlog.NewZapLogger(zlog.STDInstance()),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
		"dayu_trace_id", dayu_trace.TraceID(),
	)
	//time.Sleep(time.Second) // wait for loading task_configs
	return
}

func QueryDetail(ctx context.Context, req *QueryDetailReq, bc *conf.Bootstrap) (*v1.QueryDetailReply, error) {

	lines := req.PageSize
	if lines <= 0 || lines > 100 {
		lines = 20
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	offset := (req.Page - 1) * lines
	from, _ := utils.ParseDateTimeInSec(req.FromTime)
	to, _ := utils.ParseDateTimeInSec(req.ToTime)
	query := "traceID: " + req.TraceId + " and proxy_text-correction"

	logReq := &sls.GetLogRequest{
		From:    from,
		To:      to,
		Lines:   lines,
		Offset:  offset,
		Reverse: false,
		Query:   query,
	}
	slsC := NewSlsClient(bc.Data)
	logResp, err := slsC.Client.GetLogsV2(slsC.Project, slsC.Logstore, logReq)
	if err != nil {
		return nil, err
	}
	//fmt.Println(logResp)
	appMap, originMap, functionMap := convertConfigToMap(bc.Enums)
	logDetails := make([]*v1.QueryDetailReplyItem, 0)
	for _, logMap := range logResp.Logs {

		attributeStr := logMap["attribute"]
		var attribute entity.Attribute
		_ = json.Unmarshal([]byte(attributeStr), &attribute)

		start, _ := strconv.ParseInt(logMap["start"], 10, 64)
		startTime := utils.FormatDateTime(time.UnixMicro(start))
		end, _ := strconv.ParseInt(logMap["end"], 10, 64)
		endTime := utils.FormatDateTime(time.UnixMicro(end))
		duration, _ := strconv.ParseInt(logMap["duration"], 10, 64)
		logsStr := logMap["logs"]
		var logs []entity.Log
		_ = json.Unmarshal([]byte(logsStr), &logs)

		var input, output string
		for _, l := range logs {
			if l.Name == "input" {
				input = l.LogAttribute.Key
			} else if l.Name == "output" {
				output = l.LogAttribute.Key
			}
		}
		logDetail := &v1.QueryDetailReplyItem{
			AppId:           GetOrDefault(appMap, attribute.AppId),
			DeviceId:        attribute.DeviceId,
			Origin:          GetOrDefault(originMap, attribute.Origin),
			Function:        attribute.Function,
			FunctionName:    GetOrDefault(functionMap, attribute.Function),
			FunctionVersion: attribute.FunctionVersion,
			TraceId:         logMap["traceID"],
			ParentSpanId:    logMap["parentSpanID"],
			SpanId:          logMap["spanID"],
			StartTime:       startTime,
			EndTime:         endTime,
			Start:           start,
			End:             end,
			Duration:        duration,
			Host:            logMap["host"],
			Kind:            logMap["kind"],
			Links:           logMap["links"],
			Name:            logMap["name"],
			OtlpName:        logMap["otlp.name"],
			OtlpVersion:     logMap["otlp.version"],
			StatusCode:      logMap["statusCode"],
			StatusMessage:   logMap["statusMessage"],
			TraceState:      logMap["traceState"],
			Input:           input,
			Output:          output,
			NodeName:        GetOrDefault(functionMap, attribute.Function),
		}
		fmt.Println("logDetail: ", logDetail.TraceId)
		logDetails = append(logDetails, logDetail)
	}
	return &v1.QueryDetailReply{
		List: logDetails,
	}, nil
}
