package main

import (
	"fmt"
	"unicode"
)

func main3() {
	// 定义一个字符串
	input := "hello,,"

	// 判断字符串中是否包含中文字符
	hasChinese := false
	for _, char := range input {
		if unicode.Is(unicode.Han, char) {
			hasChinese = true
			break
		}
	}

	// 输出结果
	if hasChinese {
		fmt.Println("字符串中包含中文字符。")
	} else {
		fmt.Println("字符串中不包含中文字符。")
	}
}
