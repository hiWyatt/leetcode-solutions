package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func preorderTraversal(root *TreeNode) []int {
	// 存储遍历的节点值
	res := []int{}
	// 声明用于前序递归遍历的匿名函数 无返回值
	var traversal func(node *TreeNode)
	// 初始化用于前序递归遍历的匿名函数 无返回值
	traversal = func(node *TreeNode) {
		if node == nil {
			// 递归结束条件:节点为空
			return
		}
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	// 存储遍历的节点值
	res := []int{}
	// 声明用于中序递归遍历的匿名函数 无返回值
	var traversal func(node *TreeNode)
	// 初始化用于中序递归遍历的匿名函数 无返回值
	traversal = func(node *TreeNode) {
		if node == nil {
			// 递归结束条件:节点为空
			return
		}
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

// 后序遍历
func postorderTraversal(root *TreeNode) []int {
	// 存储遍历的节点值
	res := []int{}
	// 声明用于后序递归遍历的匿名函数 无返回值
	var traversal func(node *TreeNode)
	// 初始化用于后序递归遍历的匿名函数 无返回值
	traversal = func(node *TreeNode) {
		if node == nil {
			// 递归结束条件:节点为空
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}

func main() {
	root := TreeNode{1, nil, nil}
	TreeNode1 := TreeNode{2, nil, nil}
	TreeNode2 := TreeNode{3, nil, nil}
	root.Right = &TreeNode1
	TreeNode1.Right = &TreeNode2
	fmt.Println(preorderTraversal(&root))
	fmt.Println(inorderTraversal(&root))
	fmt.Println(postorderTraversal(&root))
}
