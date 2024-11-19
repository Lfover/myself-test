package main

import (
	"fmt"
	"regexp"
)

type data struct {
	source  string
	replace string
}

func main3() {
	text := "/i/"

	// 定义正则表达式模式和替换字符串
	patterns := map[string]string{
		"\b/i/\b": "<ph>awwbc</ph>",
		"/ɛ/":     "<ph>owwdc</ph>",
		"/æ/":     "<ph>efwoof</ph>",
	}
	//allData := []data{
	//	{source: "/i/", replace: "<ph>aww/i/bc</ph>"},
	//	{source: "/i/", replace: "<ph>111/i/bc</ph>"},
	//	{source: "/ɛ/", replace: "<ph>owwdc</ph>"},
	//	{source: "/æ/", replace: "<ph>efwoof</ph>"},
	//}
	//for _, v := range allData {
	//	if strings.Contains(text, v.source) {
	//		text = strings.ReplaceAll(text, v.source, v.replace)
	//	}
	//}

	// 遍历每个模式并进行替换
	for pattern, replacement := range patterns {
		re := regexp.MustCompile(pattern)
		text = re.ReplaceAllString(text, replacement)
	}

	fmt.Println(text) // 输出替换后的文本
}
