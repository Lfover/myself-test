package main

import (
	"fmt"
	"regexp"
)

func removeHTMLTagsAndEntities(input string) string {
	// 正则表达式匹配HTML标签
	re := regexp.MustCompile(`<[^>]*>`)
	output := re.ReplaceAllString(input, "")

	// 替换&nbsp;等HTML实体为普通空格
	re = regexp.MustCompile(`&nbsp;`)
	output = re.ReplaceAllString(output, " ")

	return output
}

func main() {
	htmlContent := `\u003cp style=\"text-align: justify;\"\u003e听力文本：\u003c/p\u003e\n\u003cp style=\"text-align: justify;\"\u003eListen and choose.\u003c/p\u003e\n\u003cp style=\"text-align: justify;\"\u003eWhen does Jane get undressed? \u003c/p\u003e\n\u003cp style=\"text-align: justify;\"\u003eBoy: What time do you go to sleep? \u003c/p\u003e\n\u003cp style=\"text-align: justify;\"\u003eGirl: Well, I get undressed and go to sleep at nine o\u0026#39;clock.\u003c/p\u003e\n\u003cp style=\"text-align: justify;\"\u003eBoy: Great! Me too.\u003c/p\u003e`

	cleanText := removeHTMLTagsAndEntities(htmlContent)
	fmt.Println(cleanText)
}
