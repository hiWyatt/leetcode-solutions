package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	// 若传入空节点,则返回false
	if root == nil {
		return false
	}
	targetSum -= root.Val // 将targetSum在遍历每层的时候都减去本层节点的值
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		// 如果剩余的targetSum为0, 则正好就是符合的结果
		return true
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum) // 否则递归找
}

func main() {
	root := TreeNode{5, nil, nil}
	TreeNode1 := TreeNode{4, nil, nil}
	TreeNode2 := TreeNode{8, nil, nil}
	TreeNode3 := TreeNode{11, nil, nil}
	TreeNode4 := TreeNode{13, nil, nil}
	TreeNode5 := TreeNode{4, nil, nil}
	TreeNode6 := TreeNode{7, nil, nil}
	TreeNode7 := TreeNode{2, nil, nil}
	TreeNode8 := TreeNode{1, nil, nil}
	root.Left = &TreeNode1
	root.Right = &TreeNode2
	TreeNode1.Left = &TreeNode3
	TreeNode2.Left = &TreeNode4
	TreeNode2.Right = &TreeNode5
	TreeNode3.Left = &TreeNode6
	TreeNode3.Right = &TreeNode7
	TreeNode5.Right = &TreeNode8

	fmt.Println(hasPathSum(&root, 22))
}
