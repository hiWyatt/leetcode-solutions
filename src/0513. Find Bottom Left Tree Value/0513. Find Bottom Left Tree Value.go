package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	//初始化列表(双向链表)
	queue := list.New()
	//定义变量res记录层级
	var res int
	//根节点入队
	queue.PushBack(root)
	//遍历列表中所有节点(包括新加入的下一层节点)
	for queue.Len() > 0 {
		//保存列表长度到length
		length := queue.Len()
		//遍历一层的所有节点(遇到孩子节点则入队)
		for i := 0; i < length; i++ {
			//队列的头节点出队
			node := queue.Remove(queue.Front()).(*TreeNode)
			//更新res 记录该层的最左节点值到res
			if i == 0 {
				res = node.Val
			}
			//若遍历到的节点有左孩子,则入队
			if node.Left != nil {
				queue.PushBack(node.Left) //入队
			}
			//若遍历到的节点有右孩子,则入队
			if node.Right != nil {
				queue.PushBack(node.Right) //入队
			}
		}
	}
	return res
}
func main() {
	root := TreeNode{1, nil, nil}
	TreeNode1 := TreeNode{2, nil, nil}
	TreeNode2 := TreeNode{3, nil, nil}
	TreeNode3 := TreeNode{4, nil, nil}
	TreeNode4 := TreeNode{5, nil, nil}
	TreeNode5 := TreeNode{6, nil, nil}
	TreeNode6 := TreeNode{7, nil, nil}
	root.Left = &TreeNode1
	root.Right = &TreeNode2
	TreeNode1.Left = &TreeNode3
	TreeNode2.Left = &TreeNode4
	TreeNode2.Right = &TreeNode5
	TreeNode4.Left = &TreeNode6

	fmt.Println(findBottomLeftValue(&root))
}
