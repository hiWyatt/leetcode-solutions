package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//传入一个结点,判断以该节点为根的树是否为平衡二叉树,若是返回true,否则返回false
func isBalanced(root *TreeNode) bool {
	//若传入结点为空,则直接返回true
	if root == nil {
		return true
	}
	//若传入结点的左右子树只要有一个不是平衡二叉树,则直接返回false
	if !isBalanced(root.Left) || !isBalanced(root.Right) { //递归
		return false
	}
	//求出以传入节点为根的左右子树的最大深度
	LeftH := maxDepth(root.Left)
	RightH := maxDepth(root.Right)
	//若以传入节点为根的左右子树的最大深度差的绝对值超过1,则返回false,否则返回true
	if abs(LeftH-RightH) > 1 {
		return false
	}
	return true
}

//传入一个结点,返回以该节点为根的树的最大深度
func maxDepth(root *TreeNode) int {
	//若传入结点为空,则直接返回0
	if root == nil {
		return 0
	}
	//得到传入结点的左右子树的最大深度
	maxLR := max(maxDepth(root.Left), maxDepth(root.Right)) //递归
	//上面求出的是以传入节点为根的左右子树的最大深度,返回值要+1表示以传入节点为根的树的最大深度
	return maxLR + 1
}

//比较传入两个的数字,返回大的数字
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//返回传入数字的绝对值
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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

	fmt.Println(isBalanced(&root))
}
