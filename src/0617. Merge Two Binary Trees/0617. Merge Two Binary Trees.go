package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 合并两棵二叉树
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// 一个节点为空,直接返回另一个节点
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	//把两棵树的元素加到一起
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)    //递归
	root1.Right = mergeTrees(root1.Right, root2.Right) //递归
	return root1
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
	root1 := TreeNode{1, nil, nil}
	TreeNode1_1 := TreeNode{3, nil, nil}
	TreeNode1_2 := TreeNode{2, nil, nil}
	TreeNode1_3 := TreeNode{5, nil, nil}
	root2 := TreeNode{2, nil, nil}
	TreeNode2_1 := TreeNode{1, nil, nil}
	TreeNode2_2 := TreeNode{3, nil, nil}
	TreeNode2_3 := TreeNode{4, nil, nil}
	TreeNode2_4 := TreeNode{7, nil, nil}
	root1.Left = &TreeNode1_1
	root1.Right = &TreeNode1_2
	TreeNode1_1.Left = &TreeNode1_3
	root2.Left = &TreeNode2_1
	root2.Right = &TreeNode2_2
	TreeNode2_1.Right = &TreeNode2_3
	TreeNode2_2.Right = &TreeNode2_4
	mergeTrees(&root1, &root2)
	fmt.Println(levelOrder(&root1))
}
