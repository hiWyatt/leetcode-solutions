package main

import "fmt"

func backspaceCompare(s string, t string) bool {
	sNew := getString(s)
	tNew := getString(t)
	if sNew == tNew {
		return true
	} else {
		return false
	}
}

func getString(s string) string {
	tempArray := []string{}
	for _, item := range s {
		tempArray = append(tempArray, string(item))
	}
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(tempArray); fastIndex++ {
		if s[fastIndex] != '#' {
			tempArray[slowIndex] = tempArray[fastIndex]
			slowIndex += 1
		} else if slowIndex != 0 {
			slowIndex -= 1
		}
	}
	if slowIndex == 0 {
		return ""
	}
	var result string
	for i := 0; i < slowIndex; i++ {
		result += tempArray[i]
	}
	return result
}

func main() {
	s := "a##c"
	t := "#a#c"
	fmt.Println(backspaceCompare(s, t))
}
