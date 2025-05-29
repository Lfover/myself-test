package main

////
//// import (
////
////	"fmt"
////	"unicode"
////
//// )
////
////	func main3() {
////		// 定义一个字符串
////		input := "hello,,"
////
////		// 判断字符串中是否包含中文字符
////		hasChinese := false
////		for _, char := range input {
////			if unicode.Is(unicode.Han, char) {
////				hasChinese = true
////				break
////			}
////		}
////
////		// 输出结果
////		if hasChinese {
////			fmt.Println("字符串中包含中文字符。")
////		} else {
////			fmt.Println("字符串中不包含中文字符。")
////		}
////	}
//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//func main12() {
//	str := "TaloWL6YDixs849zaD1y_9kV6A"
//	for i, c := range str {
//		fmt.Printf("位置 %d: 字符 '%c' (Unicode: %U, 十进制: %d)\n", i, c, c, c)
//	}
//}
//
//// 12345,2个位置-》45123
//// 整形数组，数组中的所有元素向右移动K个位置，不用额外空间
//// 双指针，连续内存
//func RemoveArray(num []int, k int) []int {
//	temp := make([]int, 0, len(num))
//	n := len(num)
//	for i := n - k; i < n; i++ {
//		temp = append(temp, num[i])
//	}
//	for i := 0; i < n-k; i++ {
//		temp = append(temp, num[i])
//	}
//	return temp
//}
//
//a:=0
//defer func{
//	fmt.Printf(a)
//}(a)
//a++
//
//package main
//
//import (
//"fmt"
//"sync"
//)
//
//func worker(id int, wg *sync.WaitGroup, ch chan int) {
//	defer wg.Done()
//	defer close(ch) // 关闭 channel
//	defer fmt.Println("Worker", id, "is done (defer in worker)")
//
//	for i := 0; i < 3; i++ {
//		ch <- i
//		fmt.Println("Worker", id, "sent", i)
//	}
//}
//// 1 0  1 1  1 2
////1
//
//
//func main13() {
//	var wg sync.WaitGroup
//	ch := make(chan int)
//
//	wg.Add(1)
//	go worker(1, &wg, ch)
//
//	go func() {
//		for num := range ch {
//			fmt.Println("Received:", num)
//		}
//	}()
//
//	wg.Wait() // 等待所有 goroutine 完成
//	fmt.Println("All workers are done (defer in main)")
//}
//
////0 1 2
