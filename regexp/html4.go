package main

import (
	"fmt"
	"html"
	"regexp"
)

func removeHTMLTags1(s string) string {
	// 使用正则表达式匹配HTML标签
	re := regexp.MustCompile(`<[^>]*>`)
	// ReplaceAllString将匹配到的HTML标签替换为空字符串
	return re.ReplaceAllString(s, "")
}
func removeHTMLTagsAndEntities2(s string) string {
	// 使用正则表达式匹配HTML标签
	re := regexp.MustCompile(`<[^>]*>`)
	// 去除HTML标签
	noTags := re.ReplaceAllString(s, "")

	// 使用html.UnescapeString来解码HTML实体
	noEntities := html.UnescapeString(noTags)

	return noEntities
}

func main() {
	htmlString := "<p><img alt=\"\" height=\"150\" src=\"https://tiku-pro-cdn.speiyou.com/imgFile/cfae8ab6-85f3-4d4e-8d71-e42a5806e97d.png\" width=\"124\" /></p>\n<p>Army&#39;s Day is in&nbsp;<u>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</u>.</p>"

	// 去除HTML标签
	a := removeHTMLTagsAndEntities2(htmlString)
	fmt.Println(a)
	cleanString := removeHTMLTags1(htmlString)

	// 打印结果
	fmt.Println(cleanString)
}
