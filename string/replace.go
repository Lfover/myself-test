package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func maskAllCharacters(input string) string {
	// 使用strings.Builder来构建新的字符串，以提高性能
	var result strings.Builder

	// 遍历输入字符串的每个字符
	for _, _ = range input {
		// 将每个字符替换为星号
		result.WriteRune('*')
	}

	// 返回替换后的字符串
	return result.String()
}

func main11() {
	// 测试方法
	input := "Hello"
	masked := maskAllCharacters(input)
	fmt.Println(masked) // 输出: **************
}
func main() {
	for {
		source := rand.NewSource(time.Now().UnixNano()) // 使用当前的纳秒生成一个随机源，也就是随机种子
		ran := rand.New(source)                         // 生成一个rand
		fmt.Println(ran.Int())
		fmt.Println(ran.Int31())
		fmt.Println(ran.Intn(5))
		//const chars = "0123456789"
		//randomDigit := make([]byte, 20)
		//seed := time.Now().UnixNano()
		//
		//// 使用时间戳作为种子初始化随机数生成器
		//rand.Seed(seed)
		//for i := range randomDigit {
		//	source := rand.NewSource(time.Now().UnixNano()) // 使用当前的纳秒生成一个随机源，也就是随机种子
		//	ran := rand.New(source)
		//	randomDigit[i] = ran
		//}
		//fmt.Println(string(randomDigit))
		//time.Sleep(1 * time.Second)
	}

}
