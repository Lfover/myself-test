package main

import (
	"fmt"
	"time"
)

func main() {
	// 给定的 Unix 毫秒时间戳
	timestamp := int64(1736909430000)

	// 将毫秒时间戳转换为纳秒时间戳
	timeInNanoseconds := timestamp * int64(time.Millisecond)

	// 使用 time.Unix(0, nanoseconds) 创建一个 time.Time 对象
	t := time.Unix(0, timeInNanoseconds)

	// 定义你想要的时间格式
	const layout = "2006-01-02 15:04:05"

	// 格式化时间
	formattedTime := t.Format(layout)

	fmt.Println(formattedTime)
	s := "你是谁"
	fmt.Println(len(s))
}
