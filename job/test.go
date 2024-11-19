package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
	"github.com/tealeg/xlsx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/sync/errgroup"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
	"sync"
	"time"
)

type QueryTraceRepo interface {
	GetQueryTraceList(ctx context.Context, search *model.QueryTraceSearch) ([]*model.DwsQsAsQueryTraceSumFd, int64, error)
}

type QueryTraceUseCase struct {
	repo     QueryTraceRepo
	taskRepo TaskRepo
	poolRepo PoolRepo
	log      *log.Helper
}

func NewQueryTraceUseCase(repo QueryTraceRepo, taskRepo TaskRepo, poolRepo PoolRepo, logger log.Logger) *QueryTraceUseCase {
	return &QueryTraceUseCase{
		repo:     repo,
		taskRepo: taskRepo,
		poolRepo: poolRepo,
		log:      log.NewHelper(logger),
	}
}

func (u *QueryTraceUseCase) InitQueryTraceTasks(ctx context.Context) error {

	//todo: 获取历史数据时间范围,自然周
	//endTime := time.Now()
	//startTime := endTime.AddDate(0, 0, -7)

	//首次上线数据范围
	startTime := time.Date(2024, 1, 24, 0, 0, 0, 0, time.Local)
	endTime := time.Date(2024, 1, 24, 23, 59, 59, 0, time.Local)

	search := model.QueryTraceSearch{
		//	AppId:        conf.BizConf.AppIdPad2,
		FunctionType: "2", //固定功能入口,语音问答
		StartTime:    &startTime,
		EndTime:      &endTime,
		Sns:          conf.BizConf.TestSn,
		PageSize:     500,
	}
	queryTraces, _, _ := u.repo.GetQueryTraceList(ctx, &search)
	if len(queryTraces) == 0 {
		return nil
	}

	u.syncQueryTraceTasks1(context.TODO(), queryTraces)

	return nil
}

func (u *QueryTraceUseCase) AddQueryTraceTasks(ctx context.Context) error {

	now := time.Now()
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	startTime := endTime.AddDate(0, 0, -1)

	search := model.QueryTraceSearch{
		FunctionType: "2", //固定功能入口,语音问答
		StartTime:    &startTime,
		EndTime:      &endTime,
		Sns:          conf.BizConf.TestSn,
		PageSize:     500,
	}
	queryTraces, _, _ := u.repo.GetQueryTraceList(ctx, &search)
	if len(queryTraces) == 0 {
		return nil
	}

	u.syncQueryTraceTasks(ctx, queryTraces)

	return nil
}
func (u *QueryTraceUseCase) AddEssayTasks(ctx context.Context) error {
	return nil
}
func InitQueryDataServiceDB() (*gorm.DB, error) {
	dsn := "query_ro:ew7K3PRKEetHAzdM@tcp(querysystem.t4vbvm4gnywo0.oceanbase.aliyuncs.com)/query_system?loc=PRC&charset=utf8mb4&parseTime=True&timeout=100s"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
}

type QueryInfo struct {
	Id            int       `json:"id"`
	TraceID       string    `json:"trace_id"`
	Sn            string    `json:"sn"`
	AudioURL      string    `json:"audio_url"`
	Payload       string    `json:"payload"`
	IntentionType string    `json:"intention_type"`
	TraceURL      string    `json:"trace_url"`
	CreatedAt     time.Time `json:"created_at"`
}
type AInfo struct {
	Id            int
	TraceID       string
	Sn            string
	AudioURL      string
	Payload       string
	IntentionType string
	TraceURL      string
	CreatedAt     time.Time
}

func (u *QueryTraceUseCase) syncQueryTraceTasks(ctx context.Context, queryTraces []*model.DwsQsAsQueryTraceSumFd) error {
	var err error

	period := 1    // 固定小学
	gradeId := 1   // 固定一年级
	subjectId := 2 // 固定数学学科
	priority := 1  // 固定优先级

	eg, _ := errgroup.WithContext(ctx)
	eg.SetLimit(100)
	qDB, err := InitQueryDataServiceDB()
	if err != nil {
		fmt.Println("db链接出错, error: ", err.Error())
		return nil
	}
	//写表格
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println("B not found")
		return nil
	}
	//写trace文件
	failedIdFile, err := os.Create("./trace_id.txt")
	if err != nil {
		fmt.Println("Error creating")
		return nil
	}
	defer failedIdFile.Close()
	writer := bufio.NewWriter(failedIdFile)
	for _, _trace := range queryTraces {
		fmt.Println(_trace.QueryTraceId)
		writer.WriteString(fmt.Sprintf(" %s \n", _trace.QueryTraceId))
		//qDB.Raw(`SELECT a.trace_id,a.sn,b.audio_url,c.payload,a.intention_type,b.audio_url,a.id ,a.created_at
		//FROM query_trace a JOIN query_detail_speech_audio b ON a.trace_id = b.trace_id JOIN query_detail_speech_text c ON a.trace_id = c.trace_id
		//WHERE a.trace_id = ?`, _trace.QueryTraceId).Scan(&result)
		var result1 QueryInfo
		err = qDB.Raw(`SELECT a.trace_id AS trace_id,a.sn as sn,a.intention_type as intention_type, a.id as id,a.created_at as created_at FROM query_trace a
		WHERE a.trace_id = ?`, _trace.QueryTraceId).Scan(&result1).Error
		var result2 QueryInfo
		err = qDB.Raw(`SELECT b.audio_url as audio_url
		FROM query_detail_speech_audio b 
		WHERE b.trace_id = ?`, _trace.QueryTraceId).Scan(&result2).Error
		var result3 QueryInfo
		err = qDB.Raw(`SELECT c.payload as payload
		FROM query_detail_speech_text c 
		WHERE c.trace_id = ?`, _trace.QueryTraceId).Scan(&result3).Error

		//var mu sync.Mutex
		//mu.Lock()
		row := sheet.AddRow()
		member := &QueryInfo{
			Id:            result1.Id,
			TraceID:       result1.TraceID,
			Sn:            result1.Sn,
			AudioURL:      result2.AudioURL,
			Payload:       result3.Payload,
			IntentionType: result1.IntentionType,
			TraceURL:      "https://aiot.vdyoo.net/genie-query/xs-chat/detail?id=" + cast.ToString(result1.Id),
			CreatedAt:     result1.CreatedAt,
		}
		in := row.WriteStruct(member, -1)
		fmt.Println(in)
		time.Sleep(time.Second * 1)
		//mu.Unlock()
		trace := _trace
		eg.Go(func() error {
			//if trace.IntentionType == "" {
			//	return nil
			//}
			//nluDetails, _ := u.repo.GetQueryDetailSpeechNluList(ctx, &model.QueryDetailSpeechNluSearch{
			//	TraceId: trace.TraceID,
			//})
			//if len(nluDetails) > 0 {
			//	//过滤非二代数据
			//	if nluDetails[0].PayLoad != "" && !strings.Contains(nluDetails[0].PayLoad, "{\"data\":[") {
			//		return nil
			//	}
			//}
			// 0. 修正优先级//todo:

			mainTskId := primitive.NewObjectID()
			taskId := primitive.NewObjectID()
			// 1. 创建主任务
			mainTsk := &model.MainTask{
				Id:       mainTskId,
				Type:     model.MainTaskTypeAudioProduce,
				Status:   model.MainStatusProcessing,
				OriginId: strconv.FormatInt(1, 10),
				Priority: priority,
				Process:  model.TaskTypeAudioProduce,
				TaskId:   taskId.Hex(),
			}
			mainInfo := &model.MainTaskInfoAudioProduce{
				ID:             1,
				AppID:          trace.AppId,
				AppName:        "二代学习机",
				Badcase:        1,
				FunctionType:   cast.ToInt64(trace.FunctionType),
				Intention:      trace.IntentionTypeName,
				VerticalDomain: u.getVerticalDomain(trace.IntentionTypeName),
				OsVersion:      trace.OsVersion,
				QueryTime:      trace.QueryTime,
				Sn:             trace.Sn,
				TraceID:        trace.QueryTraceId,
				GradeId:        gradeId,
				SubjectId:      subjectId,
				PeriodId:       period,
			}
			mainTsk.Info = mainInfo.ToMap()

			// 2. 创建子任务
			//poolId := model.GetDefaultPoolId(model.MainTaskTypeAudioProduce, period, subjectId)
			poolId := 1
			task := &model.Task{
				Id:        taskId,
				MainId:    mainTskId,
				Type:      model.TaskTypeAudioProduce,
				Status:    model.TaskStatusDefault,
				PoolId:    poolId,
				Period:    period,
				GradeId:   gradeId,
				SubjectId: subjectId,
				Info:      nil,
			}

			// 写入mongo
			_, err = u.taskRepo.CreateMainTask(ctx, mainTsk)
			if err != nil {
				u.log.WithContext(ctx).Errorf("create main task error: %v", err)
				return err
			}
			_, err = u.taskRepo.CreateTask(ctx, task, nil)
			if err != nil {
				u.log.WithContext(ctx).Errorf("create task error: %v", err)
				return err
			}

			// 3. 写入任务池
			err = u.poolRepo.Push(ctx, poolId, taskId.Hex(), model.GetPriority(mainTsk.Priority, mainTsk.CreatedAt))
			if err != nil {
				u.log.WithContext(ctx).Errorf("push task to pool error: %v", err)
				return err
			}
			return nil
		})
	}
	writer.Flush()
	err = file.Save("./planB.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if err = eg.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("结束")
	return nil
}

func (u *QueryTraceUseCase) syncQueryTraceTasks1(ctx context.Context, queryTraces []*model.DwsQsAsQueryTraceSumFd) error {
	var err error

	period := 1    // 固定小学
	gradeId := 1   // 固定一年级
	subjectId := 2 // 固定数学学科
	priority := 1  // 固定优先级

	eg, _ := errgroup.WithContext(ctx)
	eg.SetLimit(100)
	qDB, err := InitQueryDataServiceDB()
	if err != nil {
		fmt.Println("db链接出错, error: ", err.Error())
		return nil
	}
	//写表格
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println("B not found")
		return nil
	}
	//写trace文件
	failedIdFile, err := os.Create("./trace_id.txt")
	if err != nil {
		fmt.Println("Error creating")
		return nil
	}
	defer failedIdFile.Close()
	writer := bufio.NewWriter(failedIdFile)

	fn := "create_task" + strconv.FormatInt(time.Now().Unix(), 10) + ".log"
	fileLog, err := os.Create(fn)
	if err != nil {
		fileLog, _ = os.Open(fn)
	}
	defer fileLog.Close()

	ch := make(chan *QueryInfo, 500)
	for _, _trace := range queryTraces {
		fmt.Println(_trace.QueryTraceId)
		trace := _trace

		eg.Go(func() error {

			writer.WriteString(fmt.Sprintf(" %s \n", trace.QueryTraceId))
			member := &QueryInfo{}
			wg := sync.WaitGroup{}
			wg.Add(3)
			traceId := trace.QueryTraceId
			go func(traceId string) {
				var result1 QueryInfo
				err = qDB.Raw(`SELECT a.trace_id AS trace_id,a.sn as sn,a.intention_type as intention_type, a.id as id,a.created_at as created_at FROM query_trace a WHERE a.trace_id = ?`, traceId).Scan(&result1).Error
				if err != nil {
					str := fmt.Sprintf("Find A err:%v\n", err)
					fmt.Printf(str)
					fileLog.WriteString(str)
				}
				member.Id = result1.Id
				member.TraceID = result1.TraceID
				member.Sn = result1.Sn
				member.IntentionType = result1.IntentionType
				member.TraceURL = "https://aiot.vdyoo.net/genie-query/xs-chat/detail?id=" + cast.ToString(result1.Id)
				member.CreatedAt = result1.CreatedAt
				wg.Done()
			}(traceId)
			go func(traceId string) {
				var result2 QueryInfo
				err = qDB.Raw(`SELECT b.audio_url as audio_url FROM query_detail_speech_audio b WHERE b.trace_id = ?`, traceId).Scan(&result2).Error
				if err != nil {
					str := fmt.Sprintf("Find B err:%v\n", err)
					fmt.Printf(str)
					fileLog.WriteString(str)
				}
				member.AudioURL = result2.AudioURL
				wg.Done()
			}(traceId)
			go func(traceId string) {
				var result3 QueryInfo
				err = qDB.Raw(`SELECT c.payload as payload FROM query_detail_speech_text c WHERE c.trace_id = ?`, traceId).Scan(&result3).Error
				if err != nil {
					str := fmt.Sprintf("Find C err:%v\n", err)
					fmt.Printf(str)
					fileLog.WriteString(str)
				}
				member.Payload = result3.Payload
				wg.Done()
			}(traceId)
			//go func(traceId string) {
			//	res, err2 := u.repo.GetQueryTraceInfo(ctx, traceId)
			//	if err2 != nil {
			//		str := fmt.Sprintf("Find A err:%v\n", err)
			//		fmt.Printf(str)
			//		fileLog.WriteString(str)
			//	}
			//	member.WakeUpType = res.WakeUpType
			//	member.WakeUpTypeName = res.WakeUpTypeName
			//	member.StopType = res.StopType
			//	member.StopTypeName = res.StopTypeName
			//	wg.Done()
			//}(traceId)
			wg.Wait()

			//写入管道
			ch <- member
			fileLog.WriteString(string(member.Id))
			mainTskId := primitive.NewObjectID()
			taskId := primitive.NewObjectID()
			// 1. 创建主任务
			mainTsk := &model.MainTask{
				Id:       mainTskId,
				Type:     model.MainTaskTypeAudioProduce,
				Status:   model.MainStatusProcessing,
				OriginId: strconv.FormatInt(1, 10),
				Priority: priority,
				Process:  model.TaskTypeAudioProduce,
				TaskId:   taskId.Hex(),
			}
			mainInfo := &model.MainTaskInfoAudioProduce{
				ID:             1,
				AppID:          trace.AppId,
				AppName:        "二代学习机",
				Badcase:        1,
				FunctionType:   cast.ToInt64(trace.FunctionType),
				Intention:      trace.IntentionTypeName,
				VerticalDomain: u.getVerticalDomain(trace.IntentionTypeName),
				OsVersion:      trace.OsVersion,
				QueryTime:      trace.QueryTime,
				Sn:             trace.Sn,
				TraceID:        trace.QueryTraceId,
				GradeId:        gradeId,
				SubjectId:      subjectId,
				PeriodId:       period,
			}
			mainTsk.Info = mainInfo.ToMap()

			// 2. 创建子任务
			//poolId := model.GetDefaultPoolId(model.MainTaskTypeAudioProduce, period, subjectId)
			poolId := 1
			task := &model.Task{
				Id:        taskId,
				MainId:    mainTskId,
				Type:      model.TaskTypeAudioProduce,
				Status:    model.TaskStatusDefault,
				PoolId:    poolId,
				Period:    period,
				GradeId:   gradeId,
				SubjectId: subjectId,
				Info:      nil,
			}

			// 写入mongo
			_, err = u.taskRepo.CreateMainTask(ctx, mainTsk)
			if err != nil {
				u.log.WithContext(ctx).Errorf("create main task error: %v", err)
				return err
			}
			_, err = u.taskRepo.CreateTask(ctx, task, nil)
			if err != nil {
				u.log.WithContext(ctx).Errorf("create task error: %v", err)
				return err
			}

			// 3. 写入任务池
			err = u.poolRepo.Push(ctx, poolId, taskId.Hex(), model.GetPriority(mainTsk.Priority, mainTsk.CreatedAt))
			if err != nil {
				u.log.WithContext(ctx).Errorf("push task to pool error: %v", err)
				return err
			}
			return nil
		})

	}
	writer.Flush()

	if err = eg.Wait(); err != nil {
		fmt.Println(err)
	}
	close(ch)
	for data := range ch {
		row := sheet.AddRow()
		in := row.WriteStruct(data, -1)
		fmt.Println(in)
	}
	err = file.Save("./planB.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("结束")
	return nil
}

func (u *QueryTraceUseCase) syncQueryTraceTasks2(ctx context.Context, queryTraces []*model.DwsQsAsQueryTraceSumFd) error {
	var err error

	period := 1    // 固定小学
	gradeId := 1   // 固定一年级
	subjectId := 2 // 固定数学学科
	priority := 1  // 固定优先级

	eg, _ := errgroup.WithContext(ctx)
	eg.SetLimit(100)

	//写表格
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println("B not found")
		return nil
	}
	//写trace文件
	failedIdFile, err := os.Create("./trace_id.txt")
	if err != nil {
		fmt.Println("Error creating")
		return nil
	}
	defer failedIdFile.Close()
	writer := bufio.NewWriter(failedIdFile)

	fn := "create_task" + strconv.FormatInt(time.Now().Unix(), 10) + ".log"
	fileLog, err := os.Create(fn)
	if err != nil {
		fileLog, _ = os.Open(fn)
	}
	defer fileLog.Close()

	ch := make(chan *QueryInfo, 500)
	for _, _trace := range queryTraces {
		fmt.Println(_trace.QueryTraceId)
		trace := _trace

		eg.Go(func() error {
			writer.WriteString(fmt.Sprintf(" %s \n", trace.QueryTraceId))
			member := &QueryInfo{
				TraceID:       trace.QueryTraceId,
				Sn:            trace.Sn,
				AudioURL:      trace.AudioUrl,
				IntentionType: trace.IntentionTypeName,
				TraceURL:      trace.TraceDetailUrl,
			}

			//写入管道
			ch <- member
			fileLog.WriteString(string(member.Id))
			mainTskId := primitive.NewObjectID()
			taskId := primitive.NewObjectID()
			// 1. 创建主任务
			mainTsk := &model.MainTask{
				Id:       mainTskId,
				Type:     model.MainTaskTypeAudioProduce,
				Status:   model.MainStatusProcessing,
				OriginId: strconv.FormatInt(1, 10),
				Priority: priority,
				Process:  model.TaskTypeAudioProduce,
				TaskId:   taskId.Hex(),
			}
			mainInfo := &model.MainTaskInfoAudioProduce{
				ID:             1,
				AppID:          trace.AppId,
				AppName:        "二代学习机",
				Badcase:        1,
				FunctionType:   cast.ToInt64(trace.FunctionType),
				Intention:      trace.IntentionTypeName,
				VerticalDomain: u.getVerticalDomain(trace.IntentionTypeName),
				OsVersion:      trace.OsVersion,
				QueryTime:      trace.QueryTime,
				Sn:             trace.Sn,
				TraceID:        trace.QueryTraceId,
				GradeId:        gradeId,
				SubjectId:      subjectId,
				PeriodId:       period,
			}
			mainTsk.Info = mainInfo.ToMap()

			// 2. 创建子任务
			//poolId := model.GetDefaultPoolId(model.MainTaskTypeAudioProduce, period, subjectId)
			poolId := 1
			task := &model.Task{
				Id:        taskId,
				MainId:    mainTskId,
				Type:      model.TaskTypeAudioProduce,
				Status:    model.TaskStatusDefault,
				PoolId:    poolId,
				Period:    period,
				GradeId:   gradeId,
				SubjectId: subjectId,
				Info:      nil,
			}

			// 写入mongo
			_, err = u.taskRepo.CreateMainTask(ctx, mainTsk)
			if err != nil {
				u.log.WithContext(ctx).Errorf("create main task error: %v", err)
				return err
			}
			_, err = u.taskRepo.CreateTask(ctx, task, nil)
			if err != nil {
				u.log.WithContext(ctx).Errorf("create task error: %v", err)
				return err
			}

			// 3. 写入任务池
			err = u.poolRepo.Push(ctx, poolId, taskId.Hex(), model.GetPriority(mainTsk.Priority, mainTsk.CreatedAt))
			if err != nil {
				u.log.WithContext(ctx).Errorf("push task to pool error: %v", err)
				return err
			}
			return nil
		})

	}
	writer.Flush()

	if err = eg.Wait(); err != nil {
		fmt.Println(err)
	}
	close(ch)
	for data := range ch {
		row := sheet.AddRow()
		in := row.WriteStruct(data, -1)
		fmt.Println(in)
	}
	err = file.Save("./planB.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("结束")
	return nil
}

func (u *QueryTraceUseCase) getVerticalDomain(intention string) string {
	vd := ""
	if intention == "" {
		return vd
	}

	rows := conf.BizConf.DomainIntentNlu
	for _, row := range rows {
		for _, v := range row.Intent {
			if v == intention {
				vd = row.Domain
				break
			}
		}
	}

	return vd
}

func (u *QueryTraceUseCase) ClearQueryTraceTasks(ctx context.Context) error {
	return u.taskRepo.DeleteTaskData(ctx)
}
