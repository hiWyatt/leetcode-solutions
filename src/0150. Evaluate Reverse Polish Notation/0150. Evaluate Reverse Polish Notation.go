package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" {
			// 栈顶两元素相加
			value := stack[len(stack)-2] + stack[len(stack)-1]
			// 栈顶元素弹出
			stack = stack[:len(stack)-1]
			// 更新栈顶元素为计算结果
			stack[len(stack)-1] = value
		} else if tokens[i] == "-" {
			// 栈顶两元素相减
			value := stack[len(stack)-2] - stack[len(stack)-1]
			// 栈顶元素弹出
			stack = stack[:len(stack)-1]
			// 更新栈顶元素为计算结果
			stack[len(stack)-1] = value
		} else if tokens[i] == "*" {
			// 栈顶两元素相乘
			value := stack[len(stack)-2] * stack[len(stack)-1]
			// 栈顶元素弹出
			stack = stack[:len(stack)-1]
			// 更新栈顶元素为计算结果
			stack[len(stack)-1] = value
		} else if tokens[i] == "/" {
			// 栈顶两元素相除
			value := stack[len(stack)-2] / stack[len(stack)-1]
			// 栈顶元素弹出
			stack = stack[:len(stack)-1]
			// 更新栈顶元素为计算结果
			stack[len(stack)-1] = value
		} else {
			// 遇到数字,入栈即可
			value, _ := strconv.Atoi(tokens[i])
			stack = append(stack, value)
		}
	}
	return stack[0]
}
func main() {
	tokens := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(tokens))
}
