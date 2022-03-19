package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slowIndex, fastIndex := head, head
	for fastIndex != nil && fastIndex.Next != nil {
		slowIndex = slowIndex.Next
		fastIndex = fastIndex.Next.Next
		// 找到环
		if slowIndex == fastIndex {
			index1 := fastIndex
			index2 := head
			// 找环的入口
			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next
			}
			return index1
		}
	}
	return nil
}

func main() {
	head := new(ListNode)
	head.Val = 3

	node1 := new(ListNode)
	head.Next = node1
	node1.Val = 2

	node2 := new(ListNode)
	node1.Next = node2
	node2.Val = 0

	node3 := new(ListNode)
	node2.Next = node3
	node3.Val = -4

	node3.Next = node1
	fmt.Println(detectCycle(head))
}
