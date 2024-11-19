// //package main
// //
// //import (
// //	"encoding/json"
// //	"fmt"
// //	"github.com/360EntSecGroup-Skylar/excelize"
// //	"github.com/go-resty/resty/v2"
// //	"github.com/google/uuid"
// //	"github.com/spf13/cast"
// //	"golang.org/x/time/rate"
// //	"gorm.io/driver/mysql"
// //	"gorm.io/gorm"
// //	"gorm.io/gorm/logger"
// //	"time"
// //)
// //
// //type QualityControl struct {
// //	Table        string `json:"table_name"`
// //	TableId      int64  `json:"table_id"`
// //	ResourceType int    `json:"resource_type"`
// //	Type         int    `json:"type"`
// //	Msg          string `json:"msg"`
// //}
// //
// //func main() {
// //	dataDB, err := InitTestDataServiceDB()
// //	if err != nil {
// //		fmt.Println("db链接出错")
// //		return
// //	}
// //	db1, _ := dataDB.DB()
// //	defer db1.Close()
// //	//checkImage("https://static-member.tal.com/enlightenment/covers/b2598fff-d56f-4eaa-b39f-0f9c53c40e4b.png")
// //	//checkAudio("https://qz-tools.oss-cn-beijing.aliyuncs.com/jingzhunxue_tts/b5a8a6c4ed30f12fffd65b5bcae4b5f3.mp3")
// //	//checkText("我是一个好人")
// //	//业务代码
// //
// //	//filePath := "/Users/tal/Documents/test.xlsx" // 你的Excel文件路径
// //	//xlFile, err := xlsx.OpenFile(filePath)
// //	//if err != nil {
// //	//	fmt.Printf("open failed: %s\n", err)
// //	//	return
// //	//}
// //	//
// //	//for _, sheet := range xlFile.Sheets {
// //	//	for _, row := range sheet.Rows {
// //	//		if len(row.Cells) > 0 {
// //	//			id := row.Cells[0].String() // 获取第一个单元格的内容作为ID
// //	//			if cast.ToInt64(id) <= 9347 {
// //	//				continue
// //	//			}
// //	//			fmt.Printf("ID: %s\n", id)
// //	//			limiter := rate.NewLimiter(2, 2)
// //	//			time.Sleep(limiter.Reserve().Delay())
// //	//			for _, cell := range row.Cells {
// //	//				text := cell.String()
// //	//				go checkText(text, "en_dict_words", cast.ToInt64(id), dataDB)
// //	//			}
// //	//			fmt.Printf("\n")
// //	//		}
// //	//	}
// //	//}
// //
// //	xlsx, err := excelize.OpenFile("/Users/tal/Documents/test2.xlsx")
// //	if err != nil {
// //		return
// //	}
// //	rows := xlsx.GetRows("Sheet1")
// //	if len(rows) == 0 {
// //		return
// //	}
// //	for k, _ := range rows {
// //		if k == 0 {
// //			continue
// //		}
// //		//if cast.ToInt(rows[k][0]) <= 18314 {
// //		//	continue
// //		//}
// //		fmt.Println(rows[k][0])
// //		limiter := rate.NewLimiter(2, 2)
// //		time.Sleep(limiter.Reserve().Delay())
// //		if rows[k][1] != "" {
// //			go checkAudio(rows[k][1], "en_dict_words", cast.ToInt64(rows[k][0]), dataDB)
// //		}
// //		if rows[k][2] != "" {
// //			go checkAudio(rows[k][2], "en_dict_words", cast.ToInt64(rows[k][0]), dataDB)
// //		}
// //
// //		fmt.Printf("\n")
// //	}
// //}
// //
// //func checkImage(imgUrl string, table string, id int64, db *gorm.DB) {
// //	newUUID, _ := uuid.NewUUID()
// //	dataParamMap := map[string]string{
// //		"tokenId": newUUID.String(),
// //		"img":     imgUrl,
// //	}
// //	paramsMap := map[string]interface{}{
// //		"accessKey": "WLneGdDiA1YnLYlh3WSQ",
// //		"appId":     "znxx_xpad",
// //		"eventId":   "xpad_vod",
// //		"type":      "POLITY_EROTIC_VIOLENT_IMGTEXTRISK",
// //		"data":      dataParamMap,
// //	}
// //	res, err := resty.New().SetRetryCount(2).SetTimeout(20 * time.Second).R().EnableTrace().SetBody(paramsMap).Post("http://api-img-sh.fengkongcloud.com/image/v4")
// //	if err != nil {
// //		fmt.Println("image table:", table, "id:", id, "error:", err)
// //		return
// //	}
// //	var imageResp ImageResp
// //	err = json.Unmarshal(res.Body(), &imageResp)
// //	if err != nil {
// //		fmt.Println("image unmarshal table:", table, "id:", id, "error:", err)
// //		return
// //	}
// //	if imageResp.Code == 1100 {
// //		if imageResp.RiskLevel == "PASS" {
// //			fmt.Println("PASS")
// //		} else {
// //			//失败记录全部数据到一张表
// //			err := db.Table("quality_control").Create(&QualityControl{
// //				Table:        table,
// //				TableId:      id,
// //				ResourceType: 2,
// //				Type:         2,
// //				Msg:          res.String(),
// //			}).Error
// //			if err != nil {
// //				fmt.Println("save resp err table:", table, "id:", id, "error:", err)
// //				return
// //			}
// //		}
// //	} else {
// //		fmt.Println(imageResp.Message, "==========FAIL")
// //		//接口失败记录
// //		err := db.Table("quality_control").Create(&QualityControl{
// //			Table:        table,
// //			TableId:      id,
// //			ResourceType: 2,
// //			Type:         1,
// //			Msg:          res.String(),
// //		}).Error
// //		if err != nil {
// //			fmt.Println("save resp err table:", table, "id:", id, "error:", err)
// //			return
// //		}
// //	}
// //}
// //
// //type ImageResp struct {
// //	Code      int    `json:"code"`
// //	Message   string `json:"message"`
// //	RiskLevel string `json:"riskLevel"`
// //}
// //
// //func checkText(text string, table string, id int64, db *gorm.DB) {
// //	newUUID, _ := uuid.NewUUID()
// //	dataParamMap := map[string]string{
// //		"tokenId": newUUID.String(),
// //		"text":    text,
// //	}
// //	paramsMap := map[string]interface{}{
// //		"accessKey": "WLneGdDiA1YnLYlh3WSQ",
// //		"appId":     "znxx_xpad",
// //		"eventId":   "xpad_vod",
// //		"type":      "TEXTRISK",
// //		"data":      dataParamMap,
// //	}
// //	res, err := resty.New().SetRetryCount(2).SetTimeout(10 * time.Second).R().EnableTrace().SetBody(paramsMap).Post("http://api-text-sh.fengkongcloud.com/text/v4")
// //	if err != nil {
// //		fmt.Println("text table:", table, "id:", id, "error:", err)
// //		return
// //	}
// //	var textResp TextResp
// //	err = json.Unmarshal(res.Body(), &textResp)
// //	if err != nil {
// //		fmt.Println("text json unmarshal table:", table, "id:", id, "error:", err)
// //		return
// //	}
// //	if textResp.Code == 1100 {
// //		if textResp.RiskLevel == "PASS" {
// //			fmt.Println("PASS")
// //		} else {
// //			// 失败记录全部数据到一张表
// //			err := db.Table("quality_control").Create(&QualityControl{
// //				Table:        table,
// //				TableId:      id,
// //				ResourceType: 1,
// //				Type:         2,
// //				Msg:          res.String(),
// //			}).Error
// //			if err != nil {
// //				fmt.Println("save resp err table:", table, "id:", id, "error:", err)
// //				return
// //			}
// //		}
// //	} else {
// //		fmt.Println("FAIL")
// //		// 接口失败记录
// //		err := db.Table("quality_control").Create(&QualityControl{
// //			Table:        table,
// //			TableId:      id,
// //			ResourceType: 1,
// //			Type:         1,
// //			Msg:          res.String(),
// //		}).Error
// //		if err != nil {
// //			fmt.Println("save resp err table:", table, "id:", id, "error:", err)
// //			return
// //		}
// //	}
// //}
// //
// //type TextResp struct {
// //	Code      int    `json:"code"`
// //	RiskLevel string `json:"riskLevel"`
// //}
// //
// //func checkAudio(audioUrl string, table string, id int64, db *gorm.DB) {
// //	newUUID, _ := uuid.NewUUID()
// //	dataParamMap := map[string]string{}
// //	paramsMap := map[string]interface{}{
// //		"accessKey":   "WLneGdDiA1YnLYlh3WSQ",
// //		"appId":       "znxx_xpad",
// //		"eventId":     "xpad_vod",
// //		"type":        "POLITY_EROTIC_MOAN_DIRTY",
// //		"data":        dataParamMap,
// //		"contentType": "URL",
// //		"content":     audioUrl,
// //		"btId":        newUUID.String(),
// //	}
// //	res, err := resty.New().SetRetryCount(2).SetTimeout(10 * time.Second).R().EnableTrace().SetBody(paramsMap).Post("http://api-audio-sh.fengkongcloud.com/audiomessage/v4")
// //	if err != nil {
// //		fmt.Println("audio table:", table, "id:", id, "error:", err)
// //		return
// //	}
// //
// //	var audioResp AudioResp
// //	err = json.Unmarshal(res.Body(), &audioResp)
// //	if err != nil {
// //		fmt.Println("audio json unmarshal table:", table, "id:", id, "error:", err)
// //		return
// //	}
// //	if audioResp.Code == 1100 {
// //		if audioResp.AudioDetail.RiskLevel == "PASS" {
// //			//fmt.Println("PASS")
// //		} else {
// //			// 失败记录全部数据到一张表
// //			err := db.Table("quality_control").Create(&QualityControl{
// //				Table:        table,
// //				TableId:      id,
// //				ResourceType: 3,
// //				Type:         2,
// //				Msg:          res.String(),
// //			}).Error
// //			if err != nil {
// //				fmt.Println("save resp err table:", table, "id:", id, "error:", err)
// //				return
// //			}
// //		}
// //	} else {
// //		//fmt.Println("FAIL")
// //		//接口调用失败记录表
// //		err := db.Table("quality_control").Create(&QualityControl{
// //			Table:        table,
// //			TableId:      id,
// //			ResourceType: 3,
// //			Type:         1,
// //			Msg:          res.String(),
// //		}).Error
// //		if err != nil {
// //			fmt.Println("save resp err table:", table, "id:", id, "error:", err)
// //			return
// //		}
// //	}
// //}
// //
// //type AudioDetail struct {
// //	RiskLevel string `json:"riskLevel"`
// //}
// //
// //type AudioResp struct {
// //	Code        int         `json:"code"`
// //	Message     string      `json:"message"`
// //	RequestId   string      `json:"requestId"`
// //	BtId        string      `json:"btId"`
// //	AudioDetail AudioDetail `json:"detail"`
// //}
// //
// //func InitTestDataServiceDB() (*gorm.DB, error) {
// //	dsn := "taishan:E7XIzjXtQpv7mL32@tcp(rm-2zee82n4243c8o4066o.mysql.rds.aliyuncs.com:3306)/data-service?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
// //	return gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
// //}
//
// /*
// chars表
// */
// package main
//
// import (
//
//	"encoding/csv"
//	"encoding/json"
//	"fmt"
//	"github.com/360EntSecGroup-Skylar/excelize"
//	"github.com/go-resty/resty/v2"
//	"github.com/google/uuid"
//	"github.com/spf13/cast"
//	"golang.org/x/time/rate"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"gorm.io/gorm/logger"
//	"os"
//	"strings"
//	"time"
//
// )
//
//	type TextResp struct {
//		AllLabels []interface{} `json:"allLabels"`
//		AuxInfo   struct {
//			ContactResult []struct {
//				ContactString string `json:"contactString"`
//				ContactType   int    `json:"contactType"`
//			} `json:"contactResult"`
//		} `json:"auxInfo"`
//		BusinessLabels  []interface{} `json:"businessLabels"`
//		Code            int           `json:"code"`
//		FinalResult     int           `json:"finalResult"`
//		Message         string        `json:"message"`
//		RequestId       string        `json:"requestId"`
//		ResultType      int           `json:"resultType"`
//		RiskDescription string        `json:"riskDescription"`
//		RiskDetail      struct {
//		} `json:"riskDetail"`
//		RiskLabel1 string `json:"riskLabel1"`
//		RiskLabel2 string `json:"riskLabel2"`
//		RiskLabel3 string `json:"riskLabel3"`
//		RiskLevel  string `json:"riskLevel"`
//	}
//
//	type QualityControl struct {
//		Table        string `json:"table_name"`
//		TableId      int64  `json:"table_id"`
//		ResourceType int    `json:"resource_type"`
//		Type         int    `json:"type"`
//		Msg          string `json:"msg"`
//	}
//
// //拿到失败表的全部记录
//
// func write(id, url, res string) {
//
//		file, err := os.OpenFile("./chinese_chars.csv", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
//		if err != nil {
//			panic(err)
//		}
//		defer file.Close()
//
//		// 创建CSV写入器
//		writer := csv.NewWriter(file)
//		defer writer.Flush()
//
//		// 定义要写入的数据
//		data := []string{id, url, res, "chinese_chars"}
//
//		// 将数据写入CSV文件
//		if err := writer.Write(data); err != nil {
//			panic(err)
//		}
//
//		// 确保将所有数据刷新到文件
//		writer.Flush()
//
//		if err := writer.Error(); err != nil {
//			panic(err)
//		}
//	}
//
//	func getMsgAudio(msg string) string {
//		m := cast.ToStringMap(msg)
//		if detail, ok := m["detail"]; ok {
//			detailM := cast.ToStringMap(detail)
//			if audioDetail, ok := detailM["audioDetail"]; ok {
//				audioDetailList, err := cast.ToSliceE(audioDetail)
//				if err != nil {
//					return ""
//				}
//				if len(audioDetailList) != 0 {
//					audioDetailM := cast.ToStringMap(audioDetailList[0])
//					if v, ok := audioDetailM["riskDescription"].(string); ok {
//						return v
//					}
//				}
//			}
//		}
//		return ""
//	}
//
//	func main() {
//		dataDB, err := InitTestDataServiceDB()
//		if err != nil {
//			fmt.Println("db链接出错")
//			return
//		}
//
//		ubs := make([]QualityControl, 0)
//		err = dataDB.Table("quality_control").Where("table_name = ? AND type = 2 AND msg LIKE ? ", "chinese_chars", "%1901%").Find(&ubs).Error
//		if err != nil {
//			return
//		}
//
//		xlsx, err := excelize.OpenFile("/Users/tal/Documents/test3.xlsx")
//		if err != nil {
//			return
//		}
//		rows := xlsx.GetRows("Sheet1")
//		if len(rows) == 0 {
//			return
//		}
//		//
//		//dataDB2, err := InitTestDataServiceDB()
//		//if err != nil {
//		//	fmt.Println("db链接出错")
//		//	return
//		//}
//		limiter := rate.NewLimiter(2, 3)
//		for _, ub := range ubs {
//			res := &TextResp{}
//			json.Unmarshal([]byte(ub.Msg), res)
//			result := res.RiskDescription
//			time.Sleep(limiter.Reserve().Delay())
//			for k, _ := range rows {
//				if k == 0 {
//					continue
//				}
//				if cast.ToInt64(rows[k][0]) == ub.TableId {
//					fmt.Println(rows[k][0])
//					if ub.ResourceType == 1 {
//						//text := rows[k][1] + "。" + rows[k][2] + "。" + rows[k][7] + "。" + rows[k][9] + "。" + rows[k][10] + "。" + rows[k][11] + "。" + rows[k][12] + "。" + rows[k][13] + "。" + rows[k][16] + "。" + rows[k][17]
//						text := rows[k][1] + "。" + rows[k][2] + "。" + rows[k][5] + "。" + rows[k][8] + "。" + rows[k][10] + "。" + rows[k][11] + "。" + rows[k][12] + "。" + rows[k][13] + "。" + rows[k][14] + "。" + rows[k][15] + "。" + rows[k][16] + "。" + rows[k][18] + "。" + rows[k][19] + "。" + rows[k][20] + "。" + rows[k][21] + "。" + rows[k][22] + "。" + rows[k][23]
//						write(rows[k][0], text, result)
//						//go checkText(text, "chinese_chars", cast.ToInt64(rows[k][0]), dataDB2)
//					}
//					if ub.ResourceType == 2 {
//						if rows[k][6] != "" {
//							var urls3 []string
//							err = json.Unmarshal([]byte(rows[k][6]), &urls3)
//							if err != nil {
//								fmt.Println("=======1:", err)
//								return
//							}
//							for _, v := range urls3 {
//								if v == "NaN" {
//									continue
//								}
//								write(rows[k][0], v, result)
//								//fmt.Println("=======2:", v)
//								//go checkImage(v, "chinese_chars", cast.ToInt64(rows[k][0]), dataDB)
//							}
//						}
//						if rows[k][7] != "" {
//							var urls3 []string
//							err = json.Unmarshal([]byte(rows[k][7]), &urls3)
//							if err != nil {
//								fmt.Println("=======1:", err)
//								return
//							}
//							for _, v := range urls3 {
//								if v == "NaN" {
//									continue
//								}
//								//fmt.Println("=======2:", v)
//								write(rows[k][0], v, result)
//								//go checkImage(v, "chinese_chars", cast.ToInt64(rows[k][0]), dataDB)
//							}
//						}
//
//						if rows[k][17] != "" {
//							write(rows[k][0], rows[k][17], result)
//							//go checkImage(rows[k][17], "chinese_chars", cast.ToInt64(rows[k][0]), dataDB)
//						}
//					}
//					if ub.ResourceType == 3 {
//						if rows[k][3] != "" {
//							write(rows[k][0], rows[k][3], result)
//							//go checkAudio(rows[k][3], "chinese_chars", cast.ToInt64(rows[k][0]), dataDB)
//						}
//						if rows[k][4] != "" {
//							write(rows[k][0], rows[k][4], result)
//							//go checkAudio(rows[k][4], "chinese_chars", cast.ToInt64(rows[k][0]), dataDB)
//						}
//						if rows[k][9] != "" {
//							var urls3 []string
//							rows[k][9] = strings.Replace(rows[k][9], "'", "\"", -1)
//							err = json.Unmarshal([]byte(rows[k][9]), &urls3)
//							if err != nil {
//								fmt.Println("=======1:", err)
//								return
//							}
//							for _, v := range urls3 {
//								if v == "NaN" {
//									continue
//								}
//								write(rows[k][0], v, result)
//								//fmt.Println("=======2:", v)
//								//go checkAudio(v, "chinese_chars", cast.ToInt64(rows[k][0]), dataDB)
//							}
//						}
//					}
//				}
//			}
//		}
//	}
//
//	func checkImage(imgUrl string, table string, id int64, db *gorm.DB) {
//		newUUID, _ := uuid.NewUUID()
//		dataParamMap := map[string]string{
//			"tokenId": newUUID.String(),
//			"img":     imgUrl,
//		}
//		paramsMap := map[string]interface{}{
//			"accessKey": "WLneGdDiA1YnLYlh3WSQ",
//			"appId":     "znxx_xpad",
//			"eventId":   "xpad_vod",
//			"type":      "POLITY_EROTIC_VIOLENT_IMGTEXTRISK",
//			"data":      dataParamMap,
//		}
//		res, err := resty.New().SetRetryCount(2).SetTimeout(20 * time.Second).R().EnableTrace().SetBody(paramsMap).Post("http://api-img-sh.fengkongcloud.com/image/v4")
//		if err != nil {
//			fmt.Println("image table:", table, "id:", id, "error:", err)
//			return
//		}
//		var imageResp ImageResp
//		err = json.Unmarshal(res.Body(), &imageResp)
//		if err != nil {
//			fmt.Println("image unmarshal table:", table, "id:", id, "error:", err)
//			return
//		}
//		if imageResp.Code == 1100 {
//			if imageResp.RiskLevel == "PASS" {
//				fmt.Println(id, "IMAGE====PASS")
//				err := db.Table("quality_control").Where("table_id = ?", id).Updates(&QualityControl{
//					Type: 3,
//				}).Error
//				if err != nil {
//					fmt.Println("1111111save resp err table:", table, "id:", id, "error:", err)
//					return
//				}
//			} else {
//				fmt.Println(id, "IMAGE_Un_pass")
//			}
//		} else {
//			fmt.Println(id, imageResp.Message)
//		}
//	}
//
//	type ImageResp struct {
//		Code      int    `json:"code"`
//		Message   string `json:"message"`
//		RiskLevel string `json:"riskLevel"`
//	}
//
//	func checkText(text string, table string, id int64, db *gorm.DB) {
//		newUUID, _ := uuid.NewUUID()
//		dataParamMap := map[string]string{
//			"tokenId": newUUID.String(),
//			"text":    text,
//		}
//		paramsMap := map[string]interface{}{
//			"accessKey": "WLneGdDiA1YnLYlh3WSQ",
//			"appId":     "znxx_xpad",
//			"eventId":   "xpad_vod",
//			"type":      "TEXTRISK",
//			"data":      dataParamMap,
//		}
//		res, err := resty.New().SetRetryCount(2).SetTimeout(10 * time.Second).R().EnableTrace().SetBody(paramsMap).Post("http://api-text-sh.fengkongcloud.com/text/v4")
//		if err != nil {
//			fmt.Println("text table:", table, "id:", id, "error:", err)
//			return
//		}
//		var textResp TextResp
//		err = json.Unmarshal(res.Body(), &textResp)
//		if err != nil {
//			fmt.Println("text json unmarshal table:", table, "id:", id, "error:", err)
//			return
//		}
//		if textResp.Code == 1100 {
//			if textResp.RiskLevel == "PASS" {
//				fmt.Println(id, "TEXT_PASS")
//				err := db.Table("quality_control").Where("table_id = ?", id).Updates(&QualityControl{
//					Type: 3,
//				}).Error
//				if err != nil {
//					fmt.Println("22222save resp err table:", table, "id:", id, "error:", err)
//					return
//				}
//			} else {
//				fmt.Println(id, "TEXT_UN_PASS")
//			}
//		} else {
//			fmt.Println(id, textResp.Code)
//		}
//	}
//
//	func checkAudio(audioUrl string, table string, id int64, db *gorm.DB) {
//		newUUID, _ := uuid.NewUUID()
//		dataParamMap := map[string]string{}
//		paramsMap := map[string]interface{}{
//			"accessKey":   "WLneGdDiA1YnLYlh3WSQ",
//			"appId":       "znxx_xpad",
//			"eventId":     "xpad_vod",
//			"type":        "POLITY_EROTIC_MOAN_DIRTY",
//			"data":        dataParamMap,
//			"contentType": "URL",
//			"content":     audioUrl,
//			"btId":        newUUID.String(),
//		}
//		res, err := resty.New().SetRetryCount(2).SetTimeout(30 * time.Second).R().EnableTrace().SetBody(paramsMap).Post("http://api-audio-sh.fengkongcloud.com/audiomessage/v4")
//		if err != nil {
//			fmt.Println("audio table:", table, "id:", id, "error:", err)
//			return
//		}
//
//		var audioResp AudioResp
//		err = json.Unmarshal(res.Body(), &audioResp)
//		if err != nil {
//			fmt.Println("audio json unmarshal table:", table, "id:", id, "error:", err)
//			return
//		}
//		if audioResp.Code == 1100 {
//			if audioResp.AudioDetail.RiskLevel == "PASS" {
//				fmt.Println(id, "VIDEO_PASS")
//				err := db.Table("quality_control").Where("table_id = ?", id).Updates(&QualityControl{
//					Type: 3,
//				}).Error
//				if err != nil {
//					fmt.Println("33333save resp err table:", table, "id:", id, "error:", err)
//					return
//				}
//			} else {
//				fmt.Println(id, "VIDEO_Un_pass")
//			}
//		} else {
//			fmt.Println(id, audioResp.Message)
//		}
//	}
//
//	type AudioDetail struct {
//		RiskLevel string `json:"riskLevel"`
//	}
//
//	type AudioResp struct {
//		Code        int         `json:"code"`
//		Message     string      `json:"message"`
//		RequestId   string      `json:"requestId"`
//		BtId        string      `json:"btId"`
//		AudioDetail AudioDetail `json:"detail"`
//	}
//
//	func InitTestDataServiceDB() (*gorm.DB, error) {
//		dsn := "taishan:E7XIzjXtQpv7mL32@tcp(rm-2zee82n4243c8o4066o.mysql.rds.aliyuncs.com:3306)/data-service?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
//		return gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
//	}

/*
测试文件写入读取
*/
//package main
//
//import (
//	"encoding/csv"
//	"fmt"
//	"github.com/360EntSecGroup-Skylar/excelize"
//	"os"
//)
//
//func main() {
//	xlsx, err := excelize.OpenFile("/Users/tal/Documents/test2.xlsx")
//	if err != nil {
//		return
//	}
//	rows := xlsx.GetRows("Sheet1")
//	if len(rows) == 0 {
//		return
//	}
//	for k, v := range rows {
//		if k == 0 {
//			continue
//		}
//		fmt.Println(k, "-----", v)
//		fmt.Println(rows[k][1], rows[k][1])
//		break
//	}
//	write("string(1)", "1", "21")
//}
//
//func write(id, url, res string) {
//	file, err := os.OpenFile("./chinese_words.xlsx", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//	// 创建CSV写入器
//	writer := csv.NewWriter(file)
//	defer writer.Flush()
//	// 定义要写入的数据
//	data := []string{id, url, res, "chinese_chars"}
//	// 将数据写入CSV文件
//	if err := writer.Write(data); err != nil {
//		panic(err)
//	}
//	// 确保将所有数据刷新到文件
//	writer.Flush()
//	if err := writer.Error(); err != nil {
//		panic(err)
//	}
//}

//package main
//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"gorm.io/gorm/logger"
//	"runtime/debug"
//	"time"
//
//	"content-risk-checker/conf"
//	"content-risk-checker/data"
//	"content-risk-checker/model"
//	"github.com/google/uuid"
//	"github.com/spf13/cast"
//	"golang.org/x/time/rate"
//	"gorm.io/gorm"
//)
//
//var db *gorm.DB
//
//func init() {
//	conf.InitConfig(conf.ModeTest)
//	data.InitDao()
//	InitLogger()
//	db = data.Dao.Db.Debug()
//}
//
//func main() {
//	// check()
//	handleTranscoding(2, PushVideoToShumei)
//	// handleTranscoding(200, GetVideoResult)
//	// CheckXppUserBooks()
//	select {}
//}
//
//func check() {
//	go handleTranscoding(200, PushVideoToShumei)
//	go handleTranscoding(200, GetVideoResult)
//}
//
//func handleTranscoding(batch int, f func(limit int) (int, error)) {
//	defer func() {
//		if err := recover(); err != nil {
//			debug.PrintStack()
//			go handleTranscoding(batch, f)
//		}
//	}()
//
//	sleep := time.Second
//	for {
//		sleep = time.Millisecond * 100
//		numEffect, err := f(batch)
//		// 有错误或者任务量不足
//		if err != nil || numEffect < batch {
//			return
//		}
//		time.Sleep(sleep)
//	}
//}
//
//func PushVideoToShumei(batch int) (effect int, err error) {
//	logger.Info("PushVideoToShumei")
//	client := NewClient()
//	// 1. 读取数据
//	videos := make([]model.CheckVideo, 0, batch)
//	db := data.Dao.Db.Debug()
//	err = db.Model(model.CheckVideo{}).Where("status = ?", 0).Limit(batch).Find(&videos).Error
//	if err != nil {
//		panic(err)
//		return
//	}
//	effect = len(videos)
//	for _, video := range videos {
//		ctx, f := context.WithTimeout(context.Background(), time.Second*2)
//		_ = f
//		btid := uuid.New().String()
//		if len(video.Btid) > 0 {
//			btid = video.Btid
//		}
//		dt := ShuMeiReqData{
//			BtId:    btid,
//			Url:     video.VideoUrl,
//			TokenId: uuid.New().String(),
//		}
//		var resp ShuMeiRespData
//		resp, err = client.CheckVideo(ctx, dt)
//		if err != nil {
//			return
//		}
//		// 2. 更新数据
//		err = db.Model(model.CheckVideo{}).Where("id = ?", video.Id).Updates(map[string]interface{}{
//			"status":     1,
//			"btid":       btid,
//			"request_id": resp.RequestId,
//		}).Error
//	}
//	return
//}
//
//func GetVideoResult(batch int) (effect int, err error) {
//	logger.Info("PushVideoToShumei")
//	client := NewClient()
//	// 1. 读取数据
//	videos := make([]model.CheckVideo, 0, batch)
//	db := data.Dao.Db
//	err = db.Model(model.CheckVideo{}).Where("status = ?", 1).Limit(batch).Find(&videos).Error
//	if err != nil {
//		return
//	}
//	for _, video := range videos {
//		ctx := context.Background()
//		resp, err := client.GetCheckResult(ctx, video.Btid)
//		if err != nil {
//			continue
//		}
//		fmt.Println(resp)
//	}
//	return
//}
//
//func CheckXppUserBooks() {
//	page := 1
//	for {
//		checkXppUserBooks(500)
//		fmt.Println("page: ", page)
//		page++
//	}
//}
//
//func checkXppUserBooks(batch int) {
//	ubs := make([]map[string]interface{}, 0)
//	db := data.Dao.Db
//	err := db.Table("xpp_user_books").Where("status = 0").Order("id").Limit(batch).Find(&ubs).Error
//	if err != nil {
//		return
//	}
//	limiter := rate.NewLimiter(2, 3)
//	for _, ub := range ubs {
//		time.Sleep(limiter.Reserve().Delay())
//		go handleImg(ub)
//	}
//}
//
//func handleImg(ub map[string]interface{}) (err error) {
//	client := NewClient()
//	resp1, err := client.CheckImg(context.Background(), ub["cover_url"].(string))
//	if err != nil || resp1 == nil {
//		return
//	}
//	coverRzt, _ := json.Marshal(resp1)
//	riskLevel := cast.ToString(resp1["riskLevel"])
//	resp2, err := client.CheckImg(context.Background(), ub["copyright_url"].(string))
//	if err != nil || resp2 == nil {
//		return
//	}
//	copyrightRzt, _ := json.Marshal(resp2)
//	riskLevel2 := cast.ToString(resp2["riskLevel"])
//	if riskLevel2 != "PASS" {
//		riskLevel = riskLevel2
//	}
//	err = db.Table("xpp_user_books").Where("id = ?", ub["id"]).Updates(map[string]interface{}{
//		"status":                 1,
//		"risk_level":             riskLevel,
//		"cover_check_result":     coverRzt,
//		"copyright_check_result": copyrightRzt,
//	}).Error
//	logger.Info("update xpp_user_books")
//	return
////}
//
//package main
//
//import (
//	"fmt"
//	"github.com/go-resty/resty/v2"
//	"go/build"
//	"reflect"
//)
//
//
//type ControlTime struct {
//	Client *resty.Client
//}
//
//func NewControlTime(client *resty.Client) *ControlTime {
//	return &ControlTime{
//		Client: client,
//	}
//}
//
//type CommonResp struct {
//	ErrCode     int    `json:"errcode"`
//	ErrMsg      string `json:"errmsg"`
//	Trace       string `json:"trace"`
//	DataType    int    `json:"data_type"`
//	OtelTraceId string `json:"otel_trace_id"`
//	OtelSpanId  string `json:"otel_span_id"`
//}
//
//type TimeInfo struct {
//	Block        string `json:"block"`
//	Key          string `json:"key"`
//	Value        string `json:"value"`
//	DefaultValue string `json:"default_value"`
//	ExtraData    string `json:"extra_data"`
//	Desc         string `json:"desc"`
//	UpdatedAt    int64  `json:"updated_at"`
//}
//
//type DataInfo struct {
//	List  []TimeInfo `json:"list"`
//	Total int        `json:"total"`
//}
//
//type ControlTimeResp struct {
//	CommonResp
//	Data DataInfo `json:"data"`
//}
//
//const (
//	TimeControl = "/api/v1/user/config/list"
//	Key         = "literacy_video_periods_rules"
//)
//
//
//
//const code = `
//package main
//
//import "fmt"
//
//func main() {
//    fmt.Println("Hello, World!")
//}
//`
//
//func main() {
//	config := &build.Default
//	compiled, err := config.Compiler(".", []byte(code))
//	if err != nil {
//		fmt.Println("failed to build code:", err)
//		return
//	}
//
//	value := reflect.ValueOf(compiled).Elem()
//	mainFn := value.FieldByName("Main").Interface().(func() int)
//	if rc := mainFn(); rc != 0 {
//		fmt.Printf("exit code: %d\n", rc)
//	}
//}

package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/printer"
	"go/token"
	"golang.org/x/tools/go/loader"
	"reflect"
)

func main() {
	code := `
package main
import "fmt"
func main() {
    fmt.Println("Hello, dynamic Go!")
}
`
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", code, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(node.Name)

	// 构建并运行代码
	compiledCode, err := buildAndRun(node)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Output: %s\n", compiledCode)
}

func buildAndRun(node *ast.File) (string, error) {
	// 创建虚拟文件并写入代码
	fset := token.NewFileSet()
	file := &ast.File{
		Name:  &ast.Ident{Name: "main"},
		Decls: []ast.Decl{node},
	}
	f, err := build.CreateTempFile(build.Default.GOPATH, "", "main.go")
	if err != nil {
		return "", err
	}
	if err = printer.Fprint(f, fset, file); err != nil {
		return "", err
	}
	if err = f.Close(); err != nil {
		return "", err
	}

	// 编译代码
	var buildCfg build.Config
	bin, err := buildCfg.Build([]string{"-o", "main.bin", f.Name()})
	if err != nil {
		return "", err
	}
	fmt.Printf("Compiled binary: %s\n", bin)

	// 加载和运行编译的程序
	prog, err := loader.Load(&loader.Config{Build: &buildCfg}, "main.bin")
	if err != nil {
		return "", err
	}
	main, err := prog.Lookup("main.main")
	if err != nil {
		return "", err
	}
	rv := reflect.ValueOf(main).Call(nil)
	return rv[0].Interface().(string), nil
}
