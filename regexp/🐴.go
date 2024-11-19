package main

import (
	"fmt"
)

func main() {
	text := []string{
		"i",
		"ii",
		"1",
		"2",
		"3",
	}
	for i, _ := range text {
		i++
		fmt.Println(i)
		//fmt.Println(v)
	}

}
