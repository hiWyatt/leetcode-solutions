package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// check
	if root == nil {
		return root
	}
	// 相等 直接返回root节点即可
	if root == p || root == q {
		return root
	}
	// Divide
	left := lowestCommonAncestor(root.Left, p, q)   // 递归 左 用left记录找到的p或q
	right := lowestCommonAncestor(root.Right, p, q) // 递归 右 用right记录找到的p或q

	// Conquer
	// 左右两边都不为空，则根节点为祖先
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil //left和right都为空
}
func main() {
	root := TreeNode{3, nil, nil}
	TreeNode1 := TreeNode{5, nil, nil}
	TreeNode2 := TreeNode{1, nil, nil}
	TreeNode3 := TreeNode{6, nil, nil}
	TreeNode4 := TreeNode{2, nil, nil}
	TreeNode5 := TreeNode{0, nil, nil}
	TreeNode6 := TreeNode{8, nil, nil}
	TreeNode7 := TreeNode{7, nil, nil}
	TreeNode8 := TreeNode{4, nil, nil}
	root.Left = &TreeNode1
	root.Right = &TreeNode2
	TreeNode1.Left = &TreeNode3
	TreeNode1.Right = &TreeNode4
	TreeNode2.Left = &TreeNode5
	TreeNode2.Right = &TreeNode6
	TreeNode4.Left = &TreeNode7
	TreeNode4.Right = &TreeNode8
	res := lowestCommonAncestor(&root, &TreeNode1, &TreeNode8)
	fmt.Println(res)
}
