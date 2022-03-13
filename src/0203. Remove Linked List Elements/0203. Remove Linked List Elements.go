package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head
	for p := dummyHead; p != nil && p.Next != nil; {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
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
	node2.Val = 6

	node3 := new(ListNode)
	node2.Next = node3
	node3.Val = 3

	node4 := new(ListNode)
	node3.Next = node4
	node4.Val = 4

	node5 := new(ListNode)
	node4.Next = node5
	node5.Val = 5

	node6 := new(ListNode)
	node5.Next = node6
	node6.Val = 6

	val := 6
	head = removeElements(head, val)

	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}

}
