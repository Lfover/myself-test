package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func main2() {
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)

	fmt.Println(x)
}

type MyStruct struct {
	Name string
	Age  string
}

func main5() {
	str := "你好"
	eco := url.PathEscape(str)
	fmt.Println(eco)
	s := "{\"name\":\"xiaomin\",\"age\":18}"
	a := MyStruct{}
	json.Unmarshal([]byte(s), &a)
	fmt.Println(a)
}
