package main

import (
	"fmt"
	"time"
)

func main1() {
	// 定义两个时间字符串
	timeStr1 := "2024-01-10 21:14:45"
	timeStr2 := "2024-11-04 14:26:42"

	// 定义时间格式
	layout := "2006-01-02 15:04:05"

	// 将时间字符串解析为 time.Time 类型
	time1, err1 := time.Parse(layout, timeStr1)
	if err1 != nil {
		fmt.Println("解析时间1时出错:", err1)
		return
	}
	time2, err2 := time.Parse(layout, timeStr2)
	if err2 != nil {
		fmt.Println("解析时间2时出错:", err2)
		return
	}

	// 计算两个时间的差值
	duration := time2.Sub(time1)

	// 打印结果
	fmt.Printf("时间差: %v\n", duration)
}
