package main

func removeDuplicates(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) != 0 && s[i] == stack[len(stack)-1] {
			// 栈非空且出现重复项 弹栈
			stack = stack[:len(stack)-1]
		} else {
			// 栈为空或无重复项 压栈
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

func main() {
	s := "abbaca"
	s = removeDuplicates(s)
}
