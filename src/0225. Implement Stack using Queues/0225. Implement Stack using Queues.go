package main

// MyStack 栈类
type MyStack struct {
	queue []int
}

// Constructor 构造方法
func Constructor() MyStack {
	queue := make([]int, 0)
	return MyStack{queue}
}

// Push 压栈
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

// Pop 弹栈
func (this *MyStack) Pop() int {
	result := this.queue[len(this.queue)-1]
	this.queue = this.queue[:len(this.queue)-1]
	return result
}

// Top 返回栈顶元素
func (this *MyStack) Top() int {
	return this.queue[len(this.queue)-1]
}

// Empty 判断栈空
func (this *MyStack) Empty() bool {
	if len(this.queue) == 0 {
		return true
	}
	return false
}

func main() {
	myStack := Constructor()
	myStack.Push(1)
	myStack.Push(2)
	myStack.Top()   // return 2
	myStack.Pop()   // return 2
	myStack.Empty() // return False
}
