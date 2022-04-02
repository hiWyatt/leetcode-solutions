package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	// 设置节点计数器
	count := 0
	if root == nil {
		// 根节点为空
		return count
	}
	// 用切片模拟队列
	queue := make([]*TreeNode, 0)
	// 根节点入队,节点数+1
	queue = append(queue, root)
	count++
	// 遍历二叉树
	for len(queue) > 0 {
		// 记录当前层节点数
		l := len(queue)
		// 遍历一层
		for i := 0; i < l; i++ {
			// 若当前节点有子节点,则入队
			if queue[0].Left != nil {
				// 左子节点入队,节点数+1
				queue = append(queue, queue[0].Left)
				count++
			}
			if queue[0].Right != nil {
				// 右子节点入队,节点数+1
				queue = append(queue, queue[0].Right)
				count++
			}
			// 当前节点出队
			queue = queue[1:]
		}
	}
	// 返回层次遍历的节点数
	return count
}

func main() {
	root := TreeNode{1, nil, nil}
	TreeNode1 := TreeNode{2, nil, nil}
	TreeNode2 := TreeNode{3, nil, nil}
	TreeNode3 := TreeNode{4, nil, nil}
	TreeNode4 := TreeNode{5, nil, nil}
	TreeNode5 := TreeNode{6, nil, nil}
	root.Left = &TreeNode1
	root.Right = &TreeNode2
	TreeNode1.Left = &TreeNode3
	TreeNode1.Right = &TreeNode4
	TreeNode2.Left = &TreeNode5
	fmt.Println(countNodes(&root))
}
