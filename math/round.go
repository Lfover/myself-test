package main

import (
	"fmt"
	"math"
)

func main() {
	// 定义要计算的数值
	numerator := 10.0
	denominator := 3.0

	// 执行除法运算
	result := numerator / denominator
	correctness := 300
	count := 8
	// 保留两位小数
	a := math.Round(float64(correctness/count*100)) / 100
	fmt.Println(a)
	roundedResult := math.Round(result*100) / 100

	fmt.Println(roundedResult) // 输出结果，例如 3.33
}
