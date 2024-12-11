package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/xuri/excelize/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main1() {
	// 打开Excel文件
	xlFile, err := xlsx.OpenFile("example.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 遍历所有的工作表
	for _, sheet := range xlFile.Sheets {
		// 遍历工作表中的所有行
		for _, _ = range sheet.Rows {
			// 遍历行中的所有单元格

			fmt.Println()
		}
	}
}
func InitTestDataServiceDB() (*gorm.DB, error) {
	dsn := "query_rw:jcXKkjs4JLd65xDG@tcp(querysystem.t4vbvm4gnywo0.oceanbase.aliyuncs.com)/query_system?loc=PRC&charset=utf8mb4&parseTime=True&timeout=10s"
	//dsn = "lamp_rw:p6jYjyoOEsJXhOBp@tcp(rm-2zec97l8f8vf9hl2lco.mysql.rds.aliyuncs.com)/query_system?loc=PRC&charset=utf8mb4&parseTime=True&timeout=10s"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
}
func main() {
	xlsx, err := excelize.OpenFile("/Users/tal/Desktop/千问plus格式不准确异常trace.xlsx")
	if err != nil {
		return
	}

	db, err := InitTestDataServiceDB()
	if err != nil {
		fmt.Println(err)
	}
	rows, _ := xlsx.GetRows("Sheet0")
	if len(rows) == 0 {
		return
	}
	a := 0
	for k, _ := range rows {
		if k == 0 {
			continue
		}
		Id := rows[k][9]
		var cout []string
		err = db.Debug().Table("xiaosi_memory_reflective_task").Select("tal_id").Where("task_date = ?", "2024-12-08").Where("tal_id = ?", Id).Find(&cout).Error
		if err != nil {
			fmt.Println(err)
		}
		if len(cout) > 1 {
			a += 1
		}
	}
	fmt.Println(a)

}
