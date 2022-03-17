package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := new(ListNode)
	dummyNode.Next = head
	slowIndex, fastIndex := dummyNode, dummyNode
	for i := 0; fastIndex != nil; fastIndex = fastIndex.Next {
		if i > n {
			slowIndex = slowIndex.Next
		}
		i++
	}
	slowIndex.Next = slowIndex.Next.Next
	return dummyNode.Next
}

func main() {
	head := new(ListNode)
	head.Val = 1

	node1 := new(ListNode)
	head.Next = node1
	node1.Val = 2

	node2 := new(ListNode)
	node1.Next = node2
	node2.Val = 3

	node3 := new(ListNode)
	node2.Next = node3
	node3.Val = 4

	node4 := new(ListNode)
	node3.Next = node4
	node4.Val = 5

	for p := removeNthFromEnd(head, 2); p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}
