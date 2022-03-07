package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2 // 防止直接相加造成溢出
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1 // target 在右区间[mid + 1, right]
		} else {
			right = mid - 1 // target 在左区间[left, mid - 1]
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	index := search(nums, target)
	fmt.Println(index)
}
