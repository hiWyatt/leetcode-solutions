package main

// MyQueue 队列类
type MyQueue struct {
	stackIn  []int
	stackOut []int
}

// Constructor 队列类的构造方法
func Constructor() MyQueue {
	stackIn := make([]int, 0)
	stackOut := make([]int, 0)
	return MyQueue{stackIn, stackOut}
}

// Push 入队
func (this *MyQueue) Push(x int) {
	// 向输入栈添加元素
	this.stackIn = append(this.stackIn, x)
}

// Pop 出队
func (this *MyQueue) Pop() int {
	result := 0
	// 输出栈为空,将输入栈全部元素移到输出栈
	if len(this.stackOut) == 0 {
		for len(this.stackIn) != 0 {
			// 取出输入栈栈顶元素
			temp := this.stackIn[len(this.stackIn)-1]
			this.stackIn = this.stackIn[:len(this.stackIn)-1]
			// 将从输入栈取出的元素放入输出栈
			this.stackOut = append(this.stackOut, temp)
		}
	}
	// 取出输出栈栈顶元素
	result = this.stackOut[len(this.stackOut)-1]
	this.stackOut = this.stackOut[:len(this.stackOut)-1]
	return result
}

// Peek 返回队列头元素
func (this *MyQueue) Peek() int {
	result := 0
	// 输出栈为空,将输入栈全部元素移到输出栈
	if len(this.stackOut) == 0 {
		for len(this.stackIn) != 0 {
			// 取出输入栈栈顶元素
			temp := this.stackIn[len(this.stackIn)-1]
			this.stackIn = this.stackIn[:len(this.stackIn)-1]
			// 将从输入栈取出的元素放入输出栈
			this.stackOut = append(this.stackOut, temp)
		}
	}
	// 得到输出栈栈顶元素
	result = this.stackOut[len(this.stackOut)-1]
	return result
}

func (this *MyQueue) Empty() bool {
	if len(this.stackIn) == 0 && len(this.stackOut) == 0 {
		return true
	}
	return false
}

func main() {
	MyQueue := Constructor()
	MyQueue.Push(1) // queue is: [1]
	MyQueue.Push(2) // queue is: [1, 2]
	MyQueue.Peek()  // return 1
	MyQueue.Pop()   // return 1, queue is [2]
	MyQueue.Empty() // return false
}
