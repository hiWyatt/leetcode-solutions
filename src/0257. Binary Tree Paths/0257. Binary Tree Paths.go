package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	//动态创建切片 切片中的成员类型为string 当前切片具备的长度为0
	result := make([]string, 0) //结果集
	//定义函数类型变量travel
	var travel func(node *TreeNode, s string)
	//实现匿名函数并赋值给函数类型的变量travel用于递归调用
	travel = func(node *TreeNode, s string) {
		//传入的根节点其左右孩子都为空 遇到叶子节点
		if node.Left == nil && node.Right == nil {
			//拼接节点值到字符串中
			v := s + strconv.Itoa(node.Val)
			//向切片添加元素 返回新的切片result
			result = append(result, v)
			return
		}
		//非叶子节点 将该节点值拼接到字符串s后面 (中)
		s = s + strconv.Itoa(node.Val) + "->"
		//左
		if node.Left != nil {
			travel(node.Left, s) //递归
		}
		//右
		if node.Right != nil {
			travel(node.Right, s)
		}
	}
	//使用函数类型变量travel调用刚才实现的匿名函数
	travel(root, "")
	return result
}

func main() {
	root := TreeNode{1, nil, nil}
	TreeNode1 := TreeNode{2, nil, nil}
	TreeNode2 := TreeNode{3, nil, nil}
	TreeNode3 := TreeNode{5, nil, nil}
	root.Left = &TreeNode1
	root.Right = &TreeNode2
	TreeNode1.Right = &TreeNode3

	fmt.Println(binaryTreePaths(&root))
}
