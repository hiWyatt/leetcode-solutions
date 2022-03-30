package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		// 根节点为空
		return 0
	}
	// 用切片模拟队列
	queue := make([]*TreeNode, 0)
	// 根节点入队
	queue = append(queue, root)
	// 最小深度初始化为1
	minDepth := 1
	// 遍历二叉树
	for len(queue) > 0 {
		// 记录当前层节点数
		l := len(queue)
		// 遍历一层
		for i := 0; i < l; i++ {
			// 当前节点没有子节点,返回当前深度
			if queue[0].Left == nil && queue[0].Right == nil {
				return minDepth
			}
			// 若当前节点有子节点,则入队
			if queue[0].Left != nil {
				// 左子节点入队
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				// 右子节点入队
				queue = append(queue, queue[0].Right)
			}
			// 当前节点出队
			queue = queue[1:]
		}
		// 该层节点都有子节点,最小深度+1
		minDepth++
	}
	return minDepth
}

func main() {
	root := TreeNode{3, nil, nil}
	TreeNode1 := TreeNode{9, nil, nil}
	TreeNode2 := TreeNode{20, nil, nil}
	TreeNode3 := TreeNode{15, nil, nil}
	TreeNode4 := TreeNode{7, nil, nil}
	root.Left = &TreeNode1
	root.Right = &TreeNode2
	TreeNode2.Left = &TreeNode3
	TreeNode2.Right = &TreeNode4
	fmt.Println(minDepth(&root))
}
