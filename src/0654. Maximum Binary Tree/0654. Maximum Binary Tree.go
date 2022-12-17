package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//构造最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	} //若传入数组为空 则直接返回nil
	//首先找到最大值
	index := findMax(nums)
	//然后构造二叉树
	root := &TreeNode{
		Val:   nums[index],
		Left:  constructMaximumBinaryTree(nums[:index]),   //左半边
		Right: constructMaximumBinaryTree(nums[index+1:]), //右半边
	}
	//最后返回根节点
	return root
}

//找到传入数组中的最大值返回下标
func findMax(nums []int) (index int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[index] {
			index = i
		}
	}
	return
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
	nums := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree(nums)
	fmt.Println(levelOrder(root))
}
