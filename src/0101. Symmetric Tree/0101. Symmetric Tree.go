package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	if root != nil {
		// 根节点非空,左右孩子节点入队
		queue = append(queue, root.Left, root.Right)
	}
	// BFS
	for len(queue) > 0 {
		if queue[0] == nil && queue[1] == nil {
			// 同层两侧对称的一对节点都为空.整个树不一定不对称,需要继续比较同层两侧对称的其余节点队
			queue = queue[2:] // 同层两侧对称的两个节点出队
			continue          // 结束本次循环,下次循环比较同层两侧对称的其余节点队
		} else if queue[0] == nil || queue[1] == nil || queue[0].Val != queue[1].Val {
			// 同层两侧对称的一对节点的值不相等.整个树一定不对称
			return false
		} else {
			// 同层两侧对称的一对节点的值相等,不一定对称,需要继续比较其余节点队是否对称
			// 对称的节点队入队
			queue = append(queue, queue[0].Left, queue[1].Right, queue[0].Right, queue[1].Left)
			// 比较完成的节点对出队
			queue = queue[2:]
		}
	}
	return true
}

func main() {
	root := TreeNode{1, nil, nil}
	TreeNode1 := TreeNode{2, nil, nil}
	TreeNode2 := TreeNode{2, nil, nil}
	TreeNode3 := TreeNode{3, nil, nil}
	TreeNode4 := TreeNode{3, nil, nil}
	root.Left = &TreeNode1
	root.Right = &TreeNode2
	TreeNode1.Right = &TreeNode3
	TreeNode2.Right = &TreeNode4
	fmt.Println(isSymmetric(&root))
}
