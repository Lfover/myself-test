package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func unescapeUnicode(input string) (string, error) {
	// strconv.Unquote 需要输入的字符串带有双引号，因此我们先加上双引号
	//quotedInput := fmt.Sprintf("\"%s\"", input)
	unescapedText, err := strconv.Unquote(input)
	if err != nil {
		return "", err
	}
	return unescapedText, nil
}
func removeHTMLTagsAndEntities1(input string) string {
	// 正则表达式匹配HTML标签
	re := regexp.MustCompile(`<[^>]*>`)
	cleaned := re.ReplaceAllString(input, "")
	maps := map[string]string{
		"&nbsp;":  " ",
		"&amp;":   "&",
		"&mdash;": "——",
		"&#39;":   "'",
		"&quot;":  "\"",
		"&rdquo;": "”",
		"&ldquo;": "“",
	}
	for entity, replacement := range maps {
		re = regexp.MustCompile(entity)
		cleaned = re.ReplaceAllString(cleaned, replacement)
	}

	return cleaned
}
func doubleEscape(input string) string {
	re := regexp.MustCompile(`([\\"])`)
	return re.ReplaceAllString(input, `\\$1`)
}

func main() {
	text := "<p><img alt=\"\" height=\"150\" src=\"https://tiku-pro-cdn.speiyou.com/imgFile/cfae8ab6-85f3-4d4e-8d71-e42a5806e97d.png\" width=\"124\" /></p> <p>Army&#39;s Day is in&nbsp;<u>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</u>.</p>"
	// 将 Unicode 转义字符转换为实际字符
	text = doubleEscape(text)
	unescapedText, err := unescapeUnicode(text)
	if err != nil {
		fmt.Println("Error unescaping text:", err)
	}
	cleanText := removeHTMLTagsAndEntities1(unescapedText)
	fmt.Println(cleanText)
}

//import (
//	"bytes"
//	"fmt"
//	"golang.org/x/net/html"
//	"strings"
//)
//
//// replaceHTMLEntities 替换常见的HTML实体为普通字符
//func replaceHTMLEntities(input string) string {
//	replacements := map[string]string{
//		"&nbsp;":  " ",
//		"&amp;":   "&",
//		"&mdash;": "——",
//		"&#39;":   "'",
//		"&quot;":  "\"",
//	}
//
//	for entity, replacement := range replacements {
//		input = strings.ReplaceAll(input, entity, replacement)
//	}
//	return input
//}
//
//// stripTags 移除HTML标签
//func stripTags(htmlStr string) string {
//	doc, err := html.Parse(strings.NewReader(htmlStr))
//	if err != nil {
//		panic("Unable to parse HTML")
//	}
//
//	var buf bytes.Buffer
//	err = renderNode(&buf, doc)
//	if err != nil {
//		panic("Unable to render node")
//	}
//
//	return buf.String()
//}
//
//// renderNode 递归遍历节点树，并将文本节点内容写入缓冲区
//func renderNode(buf *bytes.Buffer, n *html.Node) error {
//	if n.Type == html.TextNode {
//		buf.WriteString(n.Data)
//	} else if n.Type == html.ElementNode {
//		for c := n.FirstChild; c != nil; c = c.NextSibling {
//			renderNode(buf, c)
//		}
//	}
//	return nil
//}
//
//func main() {
//	input := `<p>Hello &amp; welcome to the world of <b>Go</b> programming! &mdash; &#39;Enjoy&#39;</p>`
//	cleaned := stripTags(input)
//	cleaned = replaceHTMLEntities(cleaned)
//
//	fmt.Println(cleaned) // Output: Hello & welcome to the world of Go programming! —— 'Enjoy'
//}
