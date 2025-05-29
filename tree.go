package main

import (
	"fmt"
)

// TreeNode 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 主函数：查找路径和为targetSum的所有路径
func pathSum(root *TreeNode, targetSum int) ([][]int, int) {
	var result [][]int
	var path []int

	// 深度优先搜索寻找路径
	dfs(root, targetSum, path, &result)

	return result, len(result)
}

// dfs函数：深度优先搜索
func dfs(node *TreeNode, targetSum int, path []int, result *[][]int) {
	// 空节点，直接返回
	if node == nil {
		return
	}

	// 添加当前节点到路径
	currentPath := append(path, node.Val)

	// 计算剩余需要的和
	remainSum := targetSum - node.Val

	// 如果是叶子节点且路径和等于目标和，则记录路径
	if node.Left == nil && node.Right == nil && remainSum == 0 {
		// 创建路径副本并添加到结果中
		pathCopy := make([]int, len(currentPath))
		copy(pathCopy, currentPath)
		*result = append(*result, pathCopy)
		return
	}

	// 继续搜索左右子树
	dfs(node.Left, remainSum, currentPath, result)
	dfs(node.Right, remainSum, currentPath, result)
}

// 打印所有路径和路径数量
func printPathSum(root *TreeNode, targetSum int) {
	paths, count := pathSum(root, targetSum)

	fmt.Printf("找到 %d 条路径和为 %d 的路径:\n", count, targetSum)
	for i, path := range paths {
		fmt.Printf("路径 %d: ", i+1)
		for j, val := range path {
			if j > 0 {
				fmt.Print(" -> ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}

func main() {
	// 构建示例二叉树
	//        10
	//       /  \
	//      5    12
	//     / \
	//    4   7
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 12,
		},
	}

	// 寻找和为19的路径
	printPathSum(root, 19)

	// 寻找和为22的路径
	printPathSum(root, 22)
}
