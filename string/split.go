package main

import (
	"fmt"
	"strings"
)

func main1() {
	originalString := "Hello, World!"
	// 确保字符串长度至少为2
	if len(originalString) > 1 {
		newString := originalString[1:] // 从索引1到字符串末尾
		fmt.Println(newString)          // 输出: "ello, World!"
	} else {
		fmt.Println("The string is too short.")
	}
}
func main6() {
	// 定义一个字符串
	input := "你好 世界！ hello word"

	// 使用strings.Split函数按空格切割字符串
	parts := strings.Split(input, " ")

	// 输出切割后的部分
	fmt.Println("切割后的字符串数组：", parts)
}
func main7() {
	// 定义一个字符串数组
	stringsArray := []string{"apple", "banana", "orange"}

	// 定义一个要判断的字符串
	searchString := "banana"

	// 使用strings.Contains函数判断字符串是否在数组中
	isContained := strings.Contains(stringsArray[0], searchString)

	// 输出结果
	if isContained {
		fmt.Println("字符串在数组中。")
	} else {
		fmt.Println("字符串不在数组中。")
	}
}
func main8() {
	// 定义一个字符串数组
	stringsArray := " pian apple"

	// 定义一个要判断的字符串
	searchString := "an"

	// 使用strings.ContainsAny函数判断字符串是否在数组中
	isContained := strings.ContainsAny(stringsArray, searchString)

	// 输出结果
	if isContained {
		fmt.Println("字符串在数组中。")
	} else {
		fmt.Println("字符串不在数组中。")
	}
}
func main9() {
	// 定义一个字符串数组
	//stringsArray := []string{"apple", "banana", "orange"}

	// 定义一个要判断的字符串
	searchString := "banana"

	// 使用strings.Index函数找到字符串在数组中的索引
	index := strings.Index("stringsArray", searchString)

	// 输出结果
	if index != -1 {
		fmt.Println("字符串在数组中。")
	} else {
		fmt.Println("字符串不在数组中。")
	}
}
func main2() {
	input := "这是一个测试json数据，我们需要获取json数据。"

	// 查找"json"在字符串中的位置
	index := strings.Index(input, "json")

	// 如果找到了"json"，切割字符串
	if index != -1 {
		// 切割字符串，保留"json"及其后面的字符
		jsonData := input[index+4:]

		// 输出切割后的json数据
		fmt.Println("jsonData:", jsonData)
	} else {
		// 如果没有找到"json"，打印原始字符串
		fmt.Println("没有找到json数据。")
	}
}
