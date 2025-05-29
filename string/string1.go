package main

import (
	"fmt"
	"strings"
)

/*
给你一个字符串 s ，请你反转字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

示例 1：

输入：s = "the sky is blue"
输出："blue is sky the"
示例 2：

输入：s = "  hello world  "
输出："world hello"
解释：反转后的字符串中不能存在前导空格和尾随空格。
示例 3：

输入：s = "a good   example"
输出："example good a"
解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。
*/
func replaceNumber(strByte []byte) string {
	// 查看有多少字符
	numCount, oldSize := 0, len(strByte)
	for i := 0; i < len(strByte); i++ {
		if (strByte[i] <= '9') && (strByte[i] >= '0') {
			numCount++
		}
	}
	// 增加长度
	for i := 0; i < numCount; i++ {
		strByte = append(strByte, []byte("     ")...)
	}

	tmpBytes := []byte("number")
	// 双指针从后遍历
	leftP, rightP := oldSize-1, len(strByte)-1
	for leftP < rightP {
		rightShift := 1
		// 如果是数字则加入number
		if (strByte[leftP] <= '9') && (strByte[leftP] >= '0') {
			for i, tmpByte := range tmpBytes {
				strByte[rightP-len(tmpBytes)+i+1] = tmpByte
			}
			rightShift = len(tmpBytes)
		} else {
			strByte[rightP] = strByte[leftP]
		}
		// 更新指针
		rightP -= rightShift
		leftP -= 1
	}
	return string(strByte)
}

func main() {
	var strByte []byte

	fmt.Scanln(&strByte)
	replaceNumber(strByte)
	for i := 0; i < len(strByte); i++ {
		if strByte[i] <= '9' && strByte[i] >= '0' {
			inserElement := []byte{'n', 'u', 'm', 'b', 'e', 'r'}
			strByte = append(strByte[:i], append(inserElement, strByte[i+1:]...)...)
			i = i + len(inserElement) - 1
		}
	}

	fmt.Printf(string(strByte))

	s := " ni hao a "
	s1 := reverseString(s)
	s2 := reverseWords(s)
	fmt.Println(s2)
	fmt.Println(s1)
}
func reverseString(s string) string {
	char := []rune(s)
	n := len(char)

	start := 0
	for i := 0; i < n; i++ {
		if char[i] != ' ' || (i > 0 && char[i-1] != ' ') {
			char[start] = char[i]
			start++
		}
	}
	//去尾巴
	if char[start-1] == ' ' {
		start--
	}
	char = char[:start]
	reverse(char, 0, len(char)-1)
	//翻转字符串
	start1 := 0
	for i := 0; i < len(char); i++ {
		if char[i] == ' ' {
			reverse(char, start1, i-1)
			start1 = i + 1
		}
	}
	//翻转单词
	return s
}
func reverse(char []rune, i, j int) {
	for i < j {
		char[i], char[j] = char[j], char[i]
		i++
		j--
	}
}

func reverseWords(s string) string {
	// 1. 去除字符串前后的空格
	s = strings.TrimSpace(s)

	// 2. 按空格分割成单词数组
	words := strings.Fields(s)

	// 3. 反转单词数组
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// 4. 用单个空格连接单词，得到结果
	return strings.Join(words, " ")
}

func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}
	t := s + s
	return strings.Contains(t[1:len(t)-1], s)
}

func repeatedSubstringPattern1(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	j := 0
	next := make([]int, n)
	next[0] = j
	for i := 1; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	// next[n-1]  最长相同前后缀的长度
	if next[n-1] != 0 && n%(n-next[n-1]) == 0 {
		return true
	}
	return false
}
func aaa(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	j := 0
	arr := make([]int, n)
	arr[0] = j
	for i := 1; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = arr[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		arr[i] = j
	}
	if arr[n-1] > 0 && n%(n-arr[n-1]) == 0 {
		return true
	}
	return false
}
