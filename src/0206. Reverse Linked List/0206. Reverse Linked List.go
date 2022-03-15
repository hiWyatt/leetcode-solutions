package main

import "fmt"

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	pre := new(ListNode)
	pre = nil
	for cur := head; cur != nil; {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
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

	for p := reverseList(head); p != nil; p = p.Next {
		fmt.Println(p.Val)
	}

}
