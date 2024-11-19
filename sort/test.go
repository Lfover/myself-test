package main

import (
	"fmt"
	"sort"
)

// 定义结构体
type Item struct {
	Name  string
	Score float64
}

// 倒序排序切片的函数
func sortItemsByScore(items []Item) {
	// 使用sort.Slice进行排序，Sorts the slice given by index i to j in place using Score for comparison.
	sort.Slice(items, func(i, j int) bool {
		// 为了实现倒序排序，我们希望分数大的在前，所以这里用j的分数和i的分数比较
		return items[j].Score < items[i].Score
	})
}

func main() {
	// 创建结构体切片
	items := []Item{
		{"Alice", 92.3},
		{"Bob", 87.1},
		{"Charlie", 95.7},
		{"Dave", 88.4},
	}

	// 对结构体切片进行排序
	sortItemsByScore(items)

	// 打印排序后的结果
	for _, item := range items {
		fmt.Printf("%s: %.1f\n", item.Name, item.Score)
	}
}
