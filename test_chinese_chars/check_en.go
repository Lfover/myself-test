package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"golang.org/x/time/rate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitTestDataServiceDB() (*gorm.DB, error) {
	dsn := "souti_rw:p6jYjyoOEsJXhOBp@tcp(rm-2zeue9603s75870g0wo.mysql.rds.aliyuncs.com:3306)/tool_content_check?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
}

func InitProdDataServiceDB() (*gorm.DB, error) {
	dsn := "souti_rw:p6jYjyoOEsJXhOBp@tcp(rm-2zeue9603s75870g0wo.mysql.rds.aliyuncs.com:3306)/tool_content_test?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
}

func JudgeTxtFile() {
	// 初始化数据库
	testDB, err := InitTestDataServiceDB()
	if err != nil {
		fmt.Println("db链接出错, error: ", err.Error())
		return
	}
	prodDB, err := InitProdDataServiceDB()
	if err != nil {
		fmt.Println("db链接出错, error: ", err.Error())
		return
	}
	// 设置频控
	limiter := rate.NewLimiter(200, 1)

	start, retryCnt := 0, 0

	for {
		var enWords []*ChineseChars
		err = prodDB.Table("chinese_chars").Limit(1000).Offset(start).Find(&enWords).Error
		if err != nil || len(enWords) == 0 {
			retryCnt++
			if retryCnt > 10 {
				break
			}
			time.Sleep(time.Second * 10)
			continue
		}

		// 检查数据
		for _, enw := range enWords {

			enText := fmt.Sprintf("%s %s %s %s %s %s %s %s %s %s %s %s %s", enw.Name,
				enw.Meanings,
				enw.Synonyms,
				enw.Antonyms,
				enw.Zuci,
				enw.Zaoju,
				enw.Tongyin,
				enw.Xingjin,
				enw.ChildMean,
				enw.Zimi,
				enw.Tishi,
				enw.Xiehouyu,
				enw.Tag,
			)

			time.Sleep(limiter.Reserve().Delay())
			go checkText(enw.Id, enText, testDB, false)
		}
		start += 1000
	}
}

func main() {
	JudgeTxtFile()
}

func checkFinished(gdbId int, dataDb *gorm.DB) bool {
	var tmp CheckEnWords
	err := dataDb.Table("check_en_words").Where("en_id=? and status!=0", gdbId).First(&tmp).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func checkExisted(gdbId int, dataDb *gorm.DB) bool {
	var tmp CheckEnWords
	err := dataDb.Table("check_chinese_chars").Where("en_id=?", gdbId).First(&tmp).Error
	return err == nil
}

func createCheck(enId int, enText, rejectMsg string, status int, dataDB *gorm.DB) error {
	return dataDB.Table("check_chinese_chars").Create(&CheckEnWords{
		EnId:      enId,
		EnText:    enText,
		Status:    status,
		RejectMsg: rejectMsg,
	}).Error
}

func updateCheck(gdbId int, rejectMsg string, status int, dataDB *gorm.DB) error {
	return dataDB.Table("check_chinese_chars").Where("en_id=?", gdbId).Updates(map[string]interface{}{
		"status":     status,
		"reject_msg": rejectMsg,
	}).Error
}

func checkText(enId int, enText string, dataDB *gorm.DB, existed bool) {
	// 检查是否存在
	if !existed {
		existed = checkExisted(enId, dataDB)
	}
	// 字符串处理
	enText, _ = strconv.Unquote(strings.Replace(strconv.Quote(enText), `\\u`, `\u`, -1))

	// 发起校验请求
	newUUID, _ := uuid.NewUUID()
	dataParamMap := map[string]string{
		"tokenId": newUUID.String(),
		"text":    enText,
	}
	paramsMap := map[string]interface{}{
		"accessKey": "WLneGdDiA1YnLYlh3WSQ",
		"appId":     "znxx_xpad",
		"eventId":   "xpad_vod",
		"type":      "TEXTRISK",
		"data":      dataParamMap,
	}
	res, err := resty.New().SetRetryCount(2).SetTimeout(10 * time.Second).R().EnableTrace().SetBody(paramsMap).Post("http://api-text-sh-tmp.fengkongcloud.com/text/v4")
	// 请求失败，gdb_id不存在，保存为 0 待重试
	if err != nil && !existed {
		createCheck(enId, enText, "", 0, dataDB)
		return
	}
	var textResp TextResp
	err = json.Unmarshal(res.Body(), &textResp)
	// 返回数据有误，gdb_id不存在，保存为 0 待重试
	if err != nil && !existed {
		createCheck(enId, enText, "", 0, dataDB)
		return
	}
	if textResp.Code == 1100 {
		if textResp.RiskLevel == "PASS" { // status 更新为 2 已通过
			if existed { // 存在走update
				updateCheck(enId, "", 2, dataDB)
				return
			} else { // 不存在走create
				createCheck(enId, enText, "", 2, dataDB)
				return
			}
		} else { // status 更新为 1 待重审
			rejectMsg := fmt.Sprintf("%v %v %v %v %v", textResp.RiskLevel, textResp.RiskDescription, textResp.RiskLabel1, textResp.RiskLabel2, textResp.RiskLabel3)
			if existed { // 存在走update
				updateCheck(enId, rejectMsg, 1, dataDB)
				return
			} else { // 不存在走create
				createCheck(enId, enText, rejectMsg, 1, dataDB)
				return
			}
		}
	} else { // 触发频控， status设置为 0 待重试
		if !existed { // 不存在走create
			createCheck(enId, enText, "", 0, dataDB)
			return
		}
	}
}

type EnDictWords struct {
	Id           int    `json:"id"`
	FlMeanings   string `json:"fl_meanings"`
	FlPhrases    string `json:"fl_phrases"`
	FlComposite  string `json:"fl_composite"`
	EnThesaurus  string `json:"en_thesaurus"`
	EnAntonym    string `json:"en_antonym"`
	EnSentences  string `json:"en_sentences"`
	FlExplain    string `json:"fl_explain"`
	FlInflection string `json:"fl_inflection"`
	FlDerivate   string `json:"fl_derivate"`
	EnPhrases    string `json:"en_phrases"`
}
type ChineseChars struct {
	Id                     int        `json:"id"`
	Name                   string     `json:"name"`
	NameOrigin             string     `json:"name_origin"`
	Pinyin                 string     `json:"pinyin"`
	PinyinOrigin           string     `json:"pinyin_origin"`
	Fayin                  string     `json:"fayin"`
	FayinOrigin            string     `json:"fayin_origin"`
	Pindu                  string     `json:"pindu"`
	PinduOrigin            string     `json:"pindu_origin"`
	Structure              string     `json:"structure"`
	StructureOrigin        string     `json:"structure_origin"`
	Radical                string     `json:"radical"`
	RadicalOrigin          string     `json:"radical_origin"`
	StrokeCount            uint8      `json:"stroke_count"`
	StrokeCountOrigin      string     `json:"stroke_count_origin"`
	BishunImage            string     `json:"bishun_image"`
	BishunImageOrigin      string     `json:"bishun_image_origin"`
	BishunImageStill       string     `json:"bishun_image_still"`
	BishunImageStillOrigin string     `json:"bishun_image_still_origin"`
	StrokeJson             string     `json:"stroke_json"`
	StrokeAudioURL         string     `json:"stroke_audio_url"`
	Meanings               string     `json:"meanings"`
	MeaningsOrigin         string     `json:"meanings_origin"`
	Synonyms               string     `json:"synonyms"`
	SynonymsOrigin         string     `json:"synonyms_origin"`
	Antonyms               string     `json:"antonyms"`
	AntonymsOrigin         string     `json:"antonyms_origin"`
	Zuci                   string     `json:"zuci"`
	ZuciOrigin             string     `json:"zuci_origin"`
	Zaoju                  string     `json:"zaoju"`
	ZaojuOrigin            string     `json:"zaoju_origin"`
	Zaozi                  string     `json:"zaozi"`
	ZaoziOrigin            string     `json:"zaozi_origin"`
	Tongyin                string     `json:"tongyin"`
	TongyinOrigin          string     `json:"tongyin_origin"`
	Xingjin                string     `json:"xingjin"`
	XingjinOrigin          string     `json:"xingjin_origin"`
	Image                  string     `json:"image"`
	ImageOrigin            string     `json:"image_origin"`
	ChildMean              string     `json:"child_mean"`
	ChildMeanOrigin        string     `json:"child_mean_origin"`
	ChildMeanFayin         string     `json:"child_mean_fayin"`
	ChildMeanFayinOrigin   string     `json:"child_mean_fayin_origin"`
	Zimi                   string     `json:"zimi"`
	ZimiOrigin             string     `json:"zimi_origin"`
	Tishi                  string     `json:"tishi"`
	TishiOrigin            string     `json:"tishi_origin"`
	Xiehouyu               string     `json:"xiehouyu"`
	XiehouyuOrigin         string     `json:"xiehouyu_origin"`
	PositionInBook         string     `json:"position_in_book"`
	PositionInBookOrigin   string     `json:"position_in_book_origin"`
	Video                  string     `json:"video"`
	VideoOrigin            string     `json:"video_origin"`
	Frequence              uint64     `json:"frequence"`
	Tag                    string     `json:"tag"`
	TagOrigin              string     `json:"tag_origin"`
	DeleteTag              string     `json:"delete_tag"`
	UUID                   string     `json:"uuid"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              *time.Time `json:"updated_at"`
	DeletedAt              *time.Time `json:"deleted_at"`
	VideosURL              string     `json:"videos_url"`
}

type CheckEnWords struct {
	Id        int    `json:"id"`         // 主键ID
	EnId      int    `json:"en_id"`      // en字段 id
	EnText    string `json:"en_text"`    // en字段 text
	Status    int    `json:"status"`     // 0: 待重试 1: 待重审 2: 重试通过
	RejectMsg string `json:"reject_msg"` // 重审理由
}

type TextResp struct {
	Code            int    `json:"code"`
	RiskLevel       string `json:"riskLevel"`
	RiskDescription string `json:"riskDescription"`
	RiskLabel1      string `json:"riskLabel1"`
	RiskLabel2      string `json:"riskLabel2"`
	RiskLabel3      string `json:"riskLabel3"`
}
