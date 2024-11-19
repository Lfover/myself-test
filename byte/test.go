package main

import (
	"fmt"
	"unsafe"
)

type EmptyStruct struct{}

func main() {
	es := EmptyStruct{}
	es2 := EmptyStruct{}
	fmt.Printf("%p", &es)
	fmt.Printf("%p", &es2)
	fmt.Println(unsafe.Sizeof(es)) // 输出：1
}
