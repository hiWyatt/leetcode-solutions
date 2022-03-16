package main

import "fmt"

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		temp := cur.Next.Next
		cur.Next.Next = temp.Next
		temp.Next = cur.Next
		cur.Next = temp
		cur = temp.Next
	}
	return dummyHead.Next
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

	for p := swapPairs(head); p != nil; p = p.Next {
		fmt.Println(p.Val)
	}

}
