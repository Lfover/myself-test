package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("send", i)
		}
	}()
	for {
		select {
		case i := <-ch:
			fmt.Println("get", i)
		}
	}
}
