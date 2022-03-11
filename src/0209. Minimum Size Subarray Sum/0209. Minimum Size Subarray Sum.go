package main

func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	leftIndex, rightIndex := -1, -1
	minSubLen := length + 1
	sum := 0
	for leftIndex <= rightIndex {
		if sum < target {
			rightIndex++
			if rightIndex > length-1 {
				break
			}
			sum += nums[rightIndex]
		} else {
			leftIndex++
			if rightIndex-leftIndex+1 < minSubLen {
				minSubLen = rightIndex - leftIndex + 1
			}
			sum -= nums[leftIndex]
		}
	}
	if minSubLen == length+1 {
		return 0
	} else {
		return minSubLen
	}
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	println(minSubArrayLen(target, nums))
}
