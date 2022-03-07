package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right + 1
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 2
	index := searchInsert(nums, target)
	fmt.Println(index)
}
