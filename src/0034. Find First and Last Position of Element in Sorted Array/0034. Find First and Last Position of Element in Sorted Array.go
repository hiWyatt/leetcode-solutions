package main

import "fmt"

// 找左边界
func getLeftBorder(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target && (mid == 0 || nums[mid-1] < target) {
			return mid
		} else {
			right = mid - 1
		}
	}
	return -1
}

//找右边界
func getRightBorder(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target && (mid == len(nums)-1 || nums[mid+1] > target) {
			return mid
		} else {
			left = mid + 1
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	leftBorder := getLeftBorder(nums, target)
	rightBorder := getRightBorder(nums, target)
	return []int{leftBorder, rightBorder}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	border := searchRange(nums, target)
	fmt.Println(border)
}
