package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 { //终止条件，最后数组为空则可以返回
		return nil
	}
	idx := len(nums) / 2
	root := &TreeNode{Val: nums[idx]}
	// 左递归 root的左孩子接住下一层左区间的构造节点
	root.Left = sortedArrayToBST(nums[:idx])
	// 右递归 root的右孩子接住下一层右区间构造的节点
	root.Right = sortedArrayToBST(nums[idx+1:])

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
	nums := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(nums)
	fmt.Println(levelOrder(root))
}
