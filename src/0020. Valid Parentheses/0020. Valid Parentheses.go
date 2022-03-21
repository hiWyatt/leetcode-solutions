package main

func isValid(s string) bool {
	// 用字典事先存好括号的匹配规则
	hashmap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	// 用切片模拟栈
	stack := make([]byte, 0)
	if s == "" {
		// 空串有效
		return true
	}
	// 遍历字符串
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			// 遍历字符串时遇到左括号 压栈
			stack = append(stack, s[i])
		} else if len(stack) != 0 && stack[len(stack)-1] == hashmap[s[i]] {
			// 遍历字符串时遇到右括号 栈非空且栈顶元素与该右括号匹配 弹栈
			stack = stack[:len(stack)-1]
		} else {
			// 遍历字符串时遇到左括号 栈为空或栈顶元素与该右括号不匹配 证明有右括号多余
			return false
		}
	}
	if len(stack) != 0 {
		// 遍历完了字符串,但是栈非空,证明有左括号多余
		return false
	} else {
		return true
	}
}

func main() {
	s := "()"
	isValid(s) // return true
}
