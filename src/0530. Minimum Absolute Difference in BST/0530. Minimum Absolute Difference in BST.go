package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历的同时计算最小值
func getMinimumDifference(root *TreeNode) int {
	// 用于保留前一个节点的指针
	var prev *TreeNode
	// 定义一个比较大的值
	min := math.MaxInt64
	// 声明匿名函数
	var travel func(node *TreeNode)
	// 实现匿名函数
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left) // 递归 左子树
		if prev != nil && node.Val-prev.Val < min {
			min = node.Val - prev.Val // 更新最小差值
		}
		prev = node
		travel(node.Right)
	}
	// 调用匿名函数
	travel(root)
	return min
}

func main() {
	root := TreeNode{4, nil, nil}
	TreeNode1 := TreeNode{2, nil, nil}
	TreeNode2 := TreeNode{6, nil, nil}
	TreeNode3 := TreeNode{1, nil, nil}
	TreeNode4 := TreeNode{3, nil, nil}
	root.Left = &TreeNode1
	root.Right = &TreeNode2
	TreeNode1.Left = &TreeNode3
	TreeNode1.Right = &TreeNode4
	fmt.Println(getMinimumDifference(&root))
}
