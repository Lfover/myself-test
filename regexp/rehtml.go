package main

import (
	"fmt"
	"html"
	"strings"
)

// htmlEscapeMap 定义需要转义的字符及其对应的HTML实体。
var htmlEscapeMap = map[string]string{
	"&":  "&amp;",
	"\"": "&quot;",
	"'":  "&#39;", // 单引号
	"<":  "&lt;",
	">":  "&gt;",
	"[":  "&#91;",
	"]":  "&#93;",
}

// escapeHTML 将输入字符串中的特殊字符替换为HTML实体。
func escapeHTML(input string) string {
	var builder strings.Builder

	for _, ch := range input {
		str := string(ch)
		if escaped, ok := htmlEscapeMap[str]; ok {
			builder.WriteString(escaped)
		} else {
			builder.WriteString(str)
		}
	}

	return builder.String()
}

func aaa(input string) string {
	a := html.EscapeString(input)
	fmt.Println(a)
	return a
}
func main() {
	text := `Hello, [World]! "This is a test." & <HTML>`
	a := aaa(text)
	fmt.Println(a)
	escapedText := escapeHTML(text)
	fmt.Println("Original:", text)
	fmt.Println("Escaped: ", escapedText)
}
