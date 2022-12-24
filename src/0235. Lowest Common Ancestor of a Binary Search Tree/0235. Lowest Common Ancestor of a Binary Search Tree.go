package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for {
		if root.Val > p.Val && root.Val > q.Val {
			// root大于p q 向左搜索
			root = root.Left // 更新root
		}
		if root.Val < p.Val && root.Val < q.Val {
			// root小于p q 向右搜索
			root = root.Right
		}
		if (root.Val-p.Val)*(root.Val-q.Val) <= 0 {
			// root在p q之间,找到最近公共祖先就是root
			return root
		}
	}
	return root
}

func main() {
	root := TreeNode{6, nil, nil}
	TreeNode1 := TreeNode{2, nil, nil}
	TreeNode2 := TreeNode{8, nil, nil}
	TreeNode3 := TreeNode{0, nil, nil}
	TreeNode4 := TreeNode{4, nil, nil}
	TreeNode5 := TreeNode{7, nil, nil}
	TreeNode6 := TreeNode{9, nil, nil}
	TreeNode7 := TreeNode{3, nil, nil}
	TreeNode8 := TreeNode{5, nil, nil}
	root.Left = &TreeNode1
	root.Right = &TreeNode2
	TreeNode1.Left = &TreeNode3
	TreeNode1.Right = &TreeNode4
	TreeNode2.Left = &TreeNode5
	TreeNode2.Right = &TreeNode6
	TreeNode4.Left = &TreeNode7
	TreeNode4.Right = &TreeNode8
	res := lowestCommonAncestor(&root, &TreeNode1, &TreeNode4)
	fmt.Println(res.Val)
}
