package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	nums := []int{-9, 2, 3, 4, 5}
	//target := 3
	//fmt.Println(FindN(nums,target))
	//fmt.Println(RemoveElement(nums,target))
	arr := SortedSquares(nums)
	for _, v := range arr {
		fmt.Println(v)
	}
	//fmt.Println(MinSubArrayLen(target,nums))
}

func FindN(n []int, target int) int {
	//二分查找目标值
	left, right := 0, len(n)-1
	for left <= right {
		mid := (left + right) / 2
		if n[mid] == target {
			return mid
		}
		if n[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 移除元素
func RemoveElement(n []int, target int) int {
	left, right := 0, len(n)-1
	for left <= right {
		if n[left] == target {
			n[left] = n[right]
			right--
		} else {
			left++
		}
	}
	return left
}

// 有序数组的平方
func SortedSquares(n []int) []int {
	left, right := 0, len(n)-1
	result := make([]int, len(n))
	tar := right
	for left <= right {
		if n[left]*n[left] > n[right]*n[right] {
			result[tar] = n[left] * n[left]
			left++
		} else {
			result[tar] = n[right] * n[right]
			right--
		}
		tar--
	}

	return result
}
