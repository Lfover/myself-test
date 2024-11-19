package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func getRandomNumbersFromArray(arr []int, count int) ([]int, error) {
	// 如果请求的数量超过数组长度，返回错误
	if count > len(arr) {
		return nil, errors.New("count can't be greater than the length of the array")
	}

	// 为了不改变原始数组，创建一个数组的副本
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)

	// 初始化随机数生成器的种子
	rand.Seed(time.Now().UnixNano())

	// 随机选择指定数量的元素
	result := make([]int, count)
	for i := 0; i < count; i++ {
		// 生成一个随机索引
		randIndex := rand.Intn(len(arrCopy) - i)
		// 将随机选中的元素添加到结果数组
		result[i] = arrCopy[randIndex]
		// 将选中的元素与数组最后一个元素交换
		arrCopy[randIndex], arrCopy[len(arrCopy)-1-i] = arrCopy[len(arrCopy)-1-i], arrCopy[randIndex]
	}

	return result, nil
}
func main() {
	arr := []int{1, 2, 3, 4, 5}
	count := 3
	result, err := getRandomNumbersFromArray(arr, count)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}
