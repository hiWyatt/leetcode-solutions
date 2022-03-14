package main

// MyLinkedList 双链表(类)
type MyLinkedList struct {
	size      int
	dummyHead *Node
	dummyTail *Node
}

// Node 节点(结构体)
type Node struct {
	value int
	pre   *Node
	next  *Node
}

// Constructor 构造方法
func Constructor() MyLinkedList {
	dummyHead := new(Node)
	dummyTail := new(Node)
	dummyHead.next = dummyTail
	dummyTail.pre = dummyHead
	return MyLinkedList{0, dummyHead, dummyTail}
}

// GetNode 获取链表index号节点
func (this *MyLinkedList) GetNode(index int) *Node {
	p := this.dummyHead.next
	if index <= (this.size >> 1) {
		// 从前遍历
		for i := 0; i < index; i++ {
			p = p.next
		}
	} else {
		// 从后遍历
		p = this.dummyTail.pre
		for i := 0; i < this.size-1-index; i++ {
			p = p.pre
		}
	}
	return p
}

// Get 获取链表index号节点的数值
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		// 索引无效
		return -1
	}
	return this.GetNode(index).value
}

// AddAtHead 在链表的最前面插入一个节点
func (this *MyLinkedList) AddAtHead(val int) {
	head := new(Node)
	head.value = val
	head.pre = this.dummyHead
	head.next = this.dummyHead.next
	this.dummyHead.next = head
	head.next.pre = head
	this.size++
}

// AddAtTail 在链表的最后面插入一个节点
func (this *MyLinkedList) AddAtTail(val int) {
	tail := new(Node)
	tail.value = val
	tail.pre = this.dummyTail.pre
	tail.next = this.dummyTail
	this.dummyTail.pre = tail
	tail.pre.next = tail
	this.size++
}

// AddAtIndex 在链表index号节点前面插入一个节点
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		// 在链表的最前面插入一个节点
		this.AddAtHead(val)
		return
	}
	if index == this.size {
		this.AddAtTail(val)
		return
	}
	if index > this.size {
		return
	}
	indexNode := this.GetNode(index)
	newNode := new(Node)
	newNode.value = val
	newNode.pre = indexNode.pre
	newNode.next = indexNode
	indexNode.pre = newNode
	newNode.pre.next = newNode
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		// 索引无效
		return
	}
	indexNode := this.GetNode(index)
	indexNode.pre.next = indexNode.next
	indexNode.next.pre = indexNode.pre
	this.size--
}

func main() {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2) //链表变为1-> 2-> 3
	linkedList.Get(1)           //返回2
	linkedList.DeleteAtIndex(1) //现在链表是1-> 3
	linkedList.Get(1)           //返回3

}
