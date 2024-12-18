package main

import "fmt"

func main() {
	testClo(1, "你好", func(s string) (error, int) {
		fmt.Println("111")
		return nil, 0
	})
}

type SkillFunc func(string) (error, int)

func testClo(a int, b string, c SkillFunc) {
	fmt.Println("222")
	c("333")

}
