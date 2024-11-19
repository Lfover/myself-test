package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

// 定义源结构体
type Source struct {
	Name  string
	Age   int
	Email string
}

// 定义目标结构体
type Destination struct {
	Name  string
	Age   int
	Email string
}

func main() {
	source := &Source{
		Name:  "John Doe",
		Age:   30,
		Email: "john.doe@example.com",
	}

	var destination Destination

	// 使用 copier.Copy 将 source 的数据复制到 destination
	err := copier.Copy(&destination, source)
	if err != nil {
		fmt.Println("Error copying data:", err)
		return
	}

	fmt.Printf("Source: %+v\n", source)
	fmt.Printf("Destination: %+v\n", destination)
}
