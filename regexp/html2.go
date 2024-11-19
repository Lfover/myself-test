package main

import (
	"fmt"
	"regexp"
	"strings"
)

func removeHTMLTags(input string) string {
	// 使用正则表达式匹配Unicode转义的HTML标签
	re := regexp.MustCompile(`<[^>]*>`)
	cleaned := re.ReplaceAllString(input, "")

	// 替换&nbsp;等HTML实体为普通空格
	re = regexp.MustCompile(`&nbsp;`)
	cleaned = re.ReplaceAllString(cleaned, " ")
	re = regexp.MustCompile(`&amp;`)
	cleaned = re.ReplaceAllString(cleaned, "&")

	re = regexp.MustCompile(`&mdash;`)
	cleaned = re.ReplaceAllString(cleaned, "——")
	re = regexp.MustCompile(`&#39;`)
	cleaned = re.ReplaceAllString(cleaned, "'")
	re = regexp.MustCompile(`&quot;`)
	cleaned = re.ReplaceAllString(cleaned, "\"")
	re = regexp.MustCompile(`\\u0026mdash;`)
	cleaned = re.ReplaceAllString(cleaned, "——")
	re = regexp.MustCompile(`\\u0026hellip;`)
	cleaned = re.ReplaceAllString(cleaned, "...")
	re = regexp.MustCompile(`\\u0026quot;`)
	cleaned = re.ReplaceAllString(cleaned, "\"")
	re = regexp.MustCompile(`\\u0026ldquo;`)
	cleaned = re.ReplaceAllString(cleaned, "“")
	re = regexp.MustCompile(`\\u0026rdquo;`)
	cleaned = re.ReplaceAllString(cleaned, "”")
	re = regexp.MustCompile(`\\u003c(.*?)\\u003e`)
	// 替换所有匹配的Unicode转义HTML标签为空字符串
	cleaned = re.ReplaceAllString(cleaned, "")
	// 去除所有连续的空白字符
	cleaned = strings.Join(strings.Fields(cleaned), " ")

	return cleaned
}

func main() {
	htmlInput := `<p>&mdash;Whose teddy bear is it? Is it Amy&#39;s? </p>
<p>&mdash;No, it isn&#39;t. It&#39;s&nbsp;<u>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</u>&nbsp;.</p> \n
["\u003cp style=\"text-align: justify;\"\u003e本题考查形容词性物主代词和名词性物主代词。句意为：这是他的笔记本，那本笔记本是她的。根据用法：有名则\u0026ldquo;形\u0026rdquo;，无名则\u0026ldquo;名\u0026rdquo;，第一个划线后有名词，需填入形容词性物主代词，his；第二个划线后无名词，需填入名词性物主代词，hers。故本题答案为B。\u003c/p\u003e"]`

	textOnly := removeHTMLTags(htmlInput)
	fmt.Println(textOnly)
}
