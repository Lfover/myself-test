package main

import (
	"fmt"
)

// 假设这是你的回调函数类型
type CallbackHandler func(data string)

// 模拟一个需要回调的异步操作
func asyncOperation(handler CallbackHandler) {
	// 在实际情况下，这里可能是一些异步操作
	go func() {
		// 模拟处理并调用回调函数
		handler("result from async operation")
	}()
}

func main() {
	// 创建一个通道用于接收回调结果
	resultChan := make(chan string)

	// 定义一个回调函数，将结果发送到通道
	callback := func(data string) {
		resultChan <- data
	}

	// 调用异步操作，传入回调函数
	asyncOperation(callback)

	// 等待并获取回调结果
	result := <-resultChan

	fmt.Println("Received result:", result)
}
