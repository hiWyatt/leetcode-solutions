package main

import "fmt"

//方法一：平方后冒泡排序
/*func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
	}
	for i := 0; i < len(nums) - 1; i++ {
		for j := 0; j < len(nums)- 1- i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}*/

//方法二：双指针法
func sortedSquares(nums []int) []int {
	length := len(nums)
	left, right := 0, length-1
	newNums := make([]int, length)
	for i := len(nums) - 1; i >= 0; i-- {
		if l, r := nums[left]*nums[left], nums[right]*nums[right]; l < r {
			newNums[i] = r
			right--
		} else {
			newNums[i] = l
			left++
		}
	}
	return newNums
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	nums = sortedSquares(nums)
	for _, i := range nums {
		fmt.Println(i)
	}
}
