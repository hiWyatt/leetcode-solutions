package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	A, B := headA, headB
	for A != B {
		if A != nil {
			A = A.Next
		} else {
			A = headB
		}
		if B != nil {
			B = B.Next
		} else {
			B = headA
		}
	}
	return A
}

func main() {
	// HEAD A
	headA := new(ListNode)
	headA.Val = 4

	nodeA1 := new(ListNode)
	headA.Next = nodeA1
	nodeA1.Val = 1

	// TOTAL
	nodeT1 := new(ListNode)
	nodeA1.Next = nodeT1
	nodeT1.Val = 8

	nodeT2 := new(ListNode)
	nodeT1.Next = nodeT2
	nodeT2.Val = 4

	nodeT3 := new(ListNode)
	nodeT2.Next = nodeT3
	nodeT3.Val = 5

	// HEAD B
	headB := new(ListNode)
	headB.Val = 5

	nodeB1 := new(ListNode)
	headB.Next = nodeB1
	nodeB1.Val = 0

	nodeB2 := new(ListNode)
	nodeB1.Next = nodeB2
	nodeB2.Val = 1

	// TOTAL
	nodeB2.Next = nodeT1

	fmt.Println(getIntersectionNode(headA, headB))

}
