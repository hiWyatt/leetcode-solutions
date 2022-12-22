package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 查找BST中的众数
func findMode(root *TreeNode) []int {
	res := make([]int, 0)           // 结果集
	count := 1                      // 计数器
	max := 1                        // 最大频率
	var prev *TreeNode              // 指向前一个元素
	var travel func(node *TreeNode) // 声明匿名变量
	travel = func(node *TreeNode) { // 实现匿名函数
		if node == nil {
			return
		}
		travel(node.Left) // 递归 左
		if prev != nil && prev.Val == node.Val {
			count++ // 重复出现 计数器+1
		} else {
			count = 1 // 没重复出现 计数器置1
		}
		if count >= max { // 达到或超过最大频率
			if count > max && len(res) > 0 {
				res = []int{node.Val} // 超过最大频率 清空结果集 只把这个元素加入到结果集中
			} else {
				res = append(res, node.Val) // 达到最大频率 把这个元素加入到结果集中即可
			}
			max = count // 更新最大频率
		}
		prev = node        // prev指向该节点
		travel(node.Right) // 递归 右
	}
	travel(root) // 调用匿名函数
	return res
}
func main() {
	root := TreeNode{1, nil, nil}
	TreeNode1 := TreeNode{2, nil, nil}
	TreeNode2 := TreeNode{2, nil, nil}
	root.Right = &TreeNode1
	TreeNode1.Left = &TreeNode2
	fmt.Println(findMode(&root))
}
