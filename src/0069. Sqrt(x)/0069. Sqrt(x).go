package main

import (
	"fmt"
)

func mySqrt(x int) int {
	left, right := 1, x/2
	if x <= 1 {
		return x
	}
	for left <= right {
		mid := left + (right-left)>>2
		if mid == x/mid {
			return mid
		} else if mid < x/mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func main() {
	x := 8
	fmt.Println(mySqrt(x))
}
