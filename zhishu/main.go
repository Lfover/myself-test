package main

import (
	"fmt"
	"io"
)

func main() {
	//质数
	for i := 2; i <= 100; i++ {
		if IsZhiShu(i) {
			fmt.Println(i)
		}
	}
}

func IsZhiShu(a int) bool {
	if a == 2 {
		return true
	}
	for i := 2; i <= a/2; i++ {
		if a%i == 0 {
			return false
		}
		io.ReadWriteCloser
	}
	return true
}
