package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//二叉搜索树中的搜索
func searchBST(root *TreeNode, val int) *TreeNode {
	//若传入节点为空或找到要搜索的值,直接返回该节点
	if root == nil || root.Val == val {
		return root
	}
	//若中间节点大于目标值就向左搜索
	if root.Val > val {
		return searchBST(root.Left, val) //递归
	}
	//若中间节点小于目标节点就向右搜索
	return searchBST(root.Right, val) //递归
}

//层次遍历 验证结果
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
	TreeNode1 := TreeNode{2, nil, nil}
	TreeNode2 := TreeNode{7, nil, nil}
	TreeNode3 := TreeNode{1, nil, nil}
	TreeNode4 := TreeNode{3, nil, nil}
	root.Left = &TreeNode1
	root.Right = &TreeNode2
	TreeNode1.Left = &TreeNode3
	TreeNode1.Right = &TreeNode4
	root1 := searchBST(&root, 2)
	fmt.Println(levelOrder(root1))
}
