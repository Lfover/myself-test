package main

import (
	"fmt"
	"sync"
)

func main2() {
	wg := &sync.WaitGroup{}
	wg.Add(1) // 假设我们只需要等待一个goroutine完成

	ch := make(chan int, 10)
	ch1 := make(chan int)

	// 启动一个goroutine来发送数据
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
			fmt.Println("send", i)
		}
		close(ch)
		wg.Done()
	}()

	// 等待goroutine完成
	wg.Wait()

	for i := 0; i < 3; i++ {
		select {
		default:

			for a := range ch {
				fmt.Println("get11111", a)
			}

		case b := <-ch1:
			fmt.Println(b)
		}
	}
}
