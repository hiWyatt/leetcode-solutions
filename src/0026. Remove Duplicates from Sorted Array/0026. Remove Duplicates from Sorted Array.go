package main

import "fmt"

func removeDuplicates(nums []int) int {
	slowIndex := 1
	for fastIndex := 1; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex-1] != nums[fastIndex] {
			nums[slowIndex] = nums[fastIndex]
			slowIndex += 1
		}
	}
	return slowIndex
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums[:]))
}
