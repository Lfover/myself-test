package main

import (
	"fmt"
	"strings"
)

func main1111() {
	s := "艮（\u000E\uF8F5瘙\u0001楁\u000F）代表山、"
	//s1 := "艮（\uF8F5瘙楁）代表山、\n"
	reportData := strings.ReplaceAll(s, " ", "")
	fmt.Println(reportData)
	str1 := []string{"apple", "banana", "cherry"}
	str2 := []string{"banana", "date", "fig", "cherry"}

	result := mergeUnique(str1, str2)
	fmt.Println("Merged and Unique Array:", result)
}

func mergeUnique(base []string, toAdd []string) []string {
	// 使用map来存储base数组中的元素，方便快速查找
	existingElements := make(map[string]bool)
	for _, item := range base {
		existingElements[item] = true
	}

	// 遍历toAdd数组，如果元素不在existingElements中，则添加到base中
	for _, item := range toAdd {
		if !existingElements[item] {
			base = append(base, item)
			existingElements[item] = true
		}
	}

	return base
}
