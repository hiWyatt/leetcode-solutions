package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//定义一个全局变量pre，用来保存cur节点的前一个节点的数值
var pre int

func convertBST(root *TreeNode) *TreeNode {
	pre = 0
	traversal(root)
	return root
}

// 遍历
func traversal(cur *TreeNode) {
	if cur == nil { // 遇空就终止
		return
	}
	traversal(cur.Right) // 右 递归
	cur.Val += pre
	pre = cur.Val
	traversal(cur.Left) // 左 递归
}

// 层先遍历 测试用
func levelOrder(root *TreeNode) [][]int {
	// 用切片存储遍历节点数值
	res := make([][]int, 0)
	if root == nil {
		// 根节点为空
		return res
	}
	// 用切片模拟队列
	queue := make([]*TreeNode, 0)
	// 根节点入队
	queue = append(queue, root)
	// 遍历二叉树
	for len(queue) > 0 {
		// 记录当前层节点数
		l := len(queue)
		// 用临时切片存储每层节点的数值
		tempSlice := make([]int, 0)
		// 遍历一层
		for i := 0; i < l; i++ {
			// 若当前节点有子节点,则入队
			if queue[0].Left != nil {
				// 左子节点入队
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				// 右子节点入队
				queue = append(queue, queue[0].Right)
			}
			// 记录当前节点值
			tempSlice = append(tempSlice, queue[0].Val)
			// 当前节点出队
			queue = queue[1:]
		}
		// 将一层的节点值存入结果集
		res = append(res, tempSlice)
	}
	return res
}
func main() {
	root := TreeNode{4, nil, nil}
	TreeNode1 := TreeNode{1, nil, nil}
	TreeNode2 := TreeNode{6, nil, nil}
	TreeNode3 := TreeNode{0, nil, nil}
	TreeNode4 := TreeNode{2, nil, nil}
	TreeNode5 := TreeNode{5, nil, nil}
	TreeNode6 := TreeNode{7, nil, nil}
	TreeNode7 := TreeNode{3, nil, nil}
	TreeNode8 := TreeNode{8, nil, nil}
	root.Left = &TreeNode1
	root.Right = &TreeNode2
	TreeNode1.Left = &TreeNode3
	TreeNode1.Right = &TreeNode4
	TreeNode2.Left = &TreeNode5
	TreeNode2.Right = &TreeNode6
	TreeNode4.Right = &TreeNode7
	TreeNode6.Right = &TreeNode8
	root1 := convertBST(&root)
	fmt.Println(levelOrder(root1))
}
