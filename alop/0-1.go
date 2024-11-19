package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Knapsack 0-1背包问题
func Knapsack(values []int, weights []int, capacity int) int {
	// 初始化dp数组
	n := len(values)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// 动态规划填表
	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			if weights[i-1] <= w {
				// 可以选择放入或不放入背包
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]]+values[i-1])
			} else {
				// 不能放入背包
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	// 返回最大价值
	return dp[n][capacity]
}

func main1() {
	values := []int{60, 100, 120}
	weights := []int{10, 20, 30}
	capacity := 50

	maxValue := Knapsack(values, weights, capacity)
	fmt.Printf("The maximum value that can be put in a knapsack of capacity %d is %d\n", capacity, maxValue)
}
