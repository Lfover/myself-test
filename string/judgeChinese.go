package main

import (
	"fmt"
	"unicode"
)

// isChineseString 判断字符串是否主要是中文
func isChineseString(s string) bool {
	count, total := 0, 0
	for _, rune := range s {
		if unicode.Is(unicode.Han, rune) {
			count++
		}
		if unicode.IsLetter(rune) {
			total++
		}
	}
	return count > total/2
}

func main() {
	testStrings := []string{"Hello 世界", "你好，世界！", "English only", "只有中文", "中英混合 Hello 世界"}

	for _, str := range testStrings {
		fmt.Printf("字符串 '%s' 是中文吗? %v\n", str, isChineseString(str))
	}
}
