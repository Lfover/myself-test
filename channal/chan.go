package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	ch := make(chan int)
	for i := 0; i < 1; i++ {
		ch <- i
		fmt.Println("send", i)
	}
	close(ch)
	for i := 0; i < 1; i++ {
		select {
		case i := <-ch:
			fmt.Println("get", i)
		}
	}
}
