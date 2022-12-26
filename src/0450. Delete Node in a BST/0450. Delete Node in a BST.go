package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 通用的二叉树的删除方式
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val { // key < root 向左搜索
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if key > root.Val { // key > root 向右搜索
		root.Right = deleteNode(root.Right, key)
		return root
	}
	// 找到目标节点
	if root.Right == nil { // 目标节点的右孩子为空
		return root.Left // 返回左孩子
	}
	if root.Left == nil { // 目标节点的左孩子为空
		return root.Right // 返回右孩子
	}
	// 目标节点的左右孩子都非空
	minNode := root.Right     // 保存目标节点的右子树到minNode中
	for minNode.Left != nil { // 找到目标节点的右子树的最左面节点
		// 保存目标节点的右子树的最左面节点到minNode中
		minNode = minNode.Left
	}
	root.Val = minNode.Val               // 目标节点的值改为其右子树的最左面节点的值
	root.Right = deleteNode1(root.Right) // 目标节点的右子树
	return root
}

// 删除目标节点的右子树最左面节点
func deleteNode1(root *TreeNode) *TreeNode {
	if root.Left == nil { // 若传入节点的左孩子为空
		pRight := root.Right // 保存传入节点的右子树到pRight中
		root.Right = nil     // 删除传入节点的右子树
		return pRight        // 返回传入节点的右子树
	}
	// 若传入节点的左孩子不为空
	root.Left = deleteNode1(root.Left) // 向左搜索
	return root
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
	root := TreeNode{2, nil, nil}
	TreeNode1 := TreeNode{1, nil, nil}
	TreeNode2 := TreeNode{7, nil, nil}
	TreeNode3 := TreeNode{5, nil, nil}
	TreeNode4 := TreeNode{9, nil, nil}
	TreeNode5 := TreeNode{4, nil, nil}
	TreeNode6 := TreeNode{6, nil, nil}
	TreeNode7 := TreeNode{8, nil, nil}
	TreeNode8 := TreeNode{10, nil, nil}
	root.Left = &TreeNode1
	root.Right = &TreeNode2
	TreeNode2.Left = &TreeNode3
	TreeNode2.Right = &TreeNode4
	TreeNode3.Left = &TreeNode5
	TreeNode3.Right = &TreeNode6
	TreeNode4.Left = &TreeNode7
	TreeNode4.Right = &TreeNode8
	root1 := deleteNode(&root, 7)
	fmt.Println(levelOrder(root1))
}
