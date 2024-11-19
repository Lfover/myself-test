package main

import (
	"fmt"
	"regexp"
)

func replaceVowels(input string) string {
	replacements := map[string]string{
		"i":  "不不",
		"ii": "你好",
		"e":  "i",
		"你":  "哦可哦",
	}
	//
	//pattern := `[`
	//for key := range replacements {
	//	pattern += key
	//}
	//pattern += `]`

	pattern := "(?:"
	for key := range replacements {
		if len(pattern) > 3 { // 忽略初始的 (?: 部分
			pattern += "|"
		}
		pattern += regexp.QuoteMeta(key)
	}
	pattern += ")"

	re := regexp.MustCompile(pattern)
	result := re.ReplaceAllStringFunc(input, func(matched string) string {
		return replacements[matched]
	})

	return result
}

//	func main() {
//		input := "Thiie 你,nihwo，随机取；  ;"
//
//		//output := replaceVowels(input)
//		isSeparator := func(c rune) bool {
//			return c == ',' || c == '，' || c == '、' || c == '；' || c == ';'
//		}
//		descList := strings.FieldsFunc(input, isSeparator)
//		for _, v := range descList {
//			fmt.Println(v)
//		}
//
// }
func main() {
	// 定义一个 map，key 是 string 类型，value 是 int 数组类型
	myMap := make(map[string][]int)
	s := "觉"
	fmt.Printf("%d\n", len(s))
	// 示例数组
	arr := []int{1, 1, 2, 3, 4, 5}

	// 使用 for 循环处理数组，并写入 map
	for _, val := range arr {
		key := fmt.Sprintf("key%d", val) // 创建 key 值，可以根据需要进行修改
		myMap[key] = append(myMap[key], val)
	}

	// 打印 map 内容
	for k, v := range myMap {
		fmt.Printf("%s: %v\n", k, v)
	}
}
