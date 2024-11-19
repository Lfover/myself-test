package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	// 打开Excel文件
	xlFile, err := xlsx.OpenFile("example.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 遍历所有的工作表
	for _, sheet := range xlFile.Sheets {
		// 遍历工作表中的所有行
		for _, row := range sheet.Rows {
			// 遍历行中的所有单元格
			for _, cell := range row.Cells {
				// 获取单元格中的文本
				text, err := cell.String()
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Printf("%s\t", text)
			}
			fmt.Println()
		}
	}
}
