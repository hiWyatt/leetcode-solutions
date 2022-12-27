package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low { //如果该节点值小于最小值，则该节点更换为该节点的右节点值，继续遍历
		right := trimBST(root.Right, low, high)
		return right //返回该节点的右子树
	}
	if root.Val > high { //如果该节点的值大于最大值，则该节点更换为该节点的左节点值，继续遍历
		left := trimBST(root.Left, low, high)
		return left
	}
	root.Left = trimBST(root.Left, low, high)   // 把下一层返回的的孩子接住
	root.Right = trimBST(root.Right, low, high) // 把下一层返回的的孩子接住
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
	root := TreeNode{3, nil, nil}
	TreeNode1 := TreeNode{0, nil, nil}
	TreeNode2 := TreeNode{4, nil, nil}
	TreeNode3 := TreeNode{2, nil, nil}
	TreeNode4 := TreeNode{1, nil, nil}
	root.Left = &TreeNode1
	root.Right = &TreeNode2
	TreeNode1.Right = &TreeNode3
	TreeNode3.Left = &TreeNode4
	root1 := trimBST(&root, 1, 3)
	fmt.Println(levelOrder(root1))
}
