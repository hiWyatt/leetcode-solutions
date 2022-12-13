package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	var res int //定义变量res 结果集
	findLeft(root, &res)
	return res
}

//递归DFS搜索到所有左叶子节点,和值计入res
func findLeft(root *TreeNode, res *int) {
	//以传入节点为根节点的 左孩子非空 且 该左孩子的左右孩子为空(是叶子节点)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		//将该左孩子的节点值加入结果集res
		*res = *res + root.Left.Val
	}
	//以传入节点为根节点的 左孩子非空 但 该左孩子的左右孩子至少一个非空(不是叶子节点)
	if root.Left != nil {
		findLeft(root.Left, res) //递归
	}
	//以传入节点为根节点的 右孩子非空 但 该右孩子的左右孩子至少一个非空(不是叶子节点)
	if root.Right != nil {
		findLeft(root.Right, res) //递归
	}
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

	fmt.Println(sumOfLeftLeaves(&root))
}
