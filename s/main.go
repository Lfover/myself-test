package main

import (
	"fmt"
)

/*
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
*/
func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		panic("defer")
	}()
	panic("main")
}

func LongestStr(s string) int {
	//哈希
	hash := make(map[byte]int)
	start := 0
	count := 0
	for i := 0; i < len(s); i++ {
		if v, ok := hash[s[i]]; ok && i > v {
			start = v + 1
		}
		hash[s[i]] = i
		currentLength := i - start + 1
		if currentLength > count {
			count = currentLength
		}
	}
	return count
}
