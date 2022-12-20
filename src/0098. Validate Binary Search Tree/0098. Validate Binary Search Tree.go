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

//判断是否为二叉搜索树
func isValidBST(root *TreeNode) bool {
	// 二叉搜索树也可以是空树
	if root == nil {
		return true
	}
	// 由题目中的数据限制可以得出min和max
	return check(root, math.MinInt64, math.MaxInt64)
}

func check(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}

	if min >= int64(node.Val) || max <= int64(node.Val) {
		return false
	}
	// 分别对左子树和右子树递归判断，如果左子树和右子树都符合则返回true
	return check(node.Right, int64(node.Val), max) && check(node.Left, min, int64(node.Val))
}
func main() {
	root := TreeNode{5, nil, nil}
	TreeNode1 := TreeNode{1, nil, nil}
	TreeNode2 := TreeNode{7, nil, nil}
	TreeNode3 := TreeNode{6, nil, nil}
	TreeNode4 := TreeNode{8, nil, nil}
	root.Left = &TreeNode1
	root.Right = &TreeNode2
	TreeNode2.Left = &TreeNode3
	TreeNode2.Right = &TreeNode4
	fmt.Println(isValidBST(&root))
}
