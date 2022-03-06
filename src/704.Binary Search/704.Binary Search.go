package main

func main() {
	nums := [...]int{-1, 0, 3, 5, 9, 12}
	target := 9
	index := search(nums, target)
	println(index)
}
func search(nums [6]int, target int) int {
	var right, left, mid int
	left = 0
	right = len(nums) - 1
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
