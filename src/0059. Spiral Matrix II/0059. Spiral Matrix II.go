package main

import "fmt"

func generateMatrix(n int) [][]int {
	left, right, top, bottom := 0, n-1, 0, n-1
	num := 1
	target := n * n
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for num <= target {
		// left to right
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		// top to bottom
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		// right to left
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		// bottom to top
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}

func main() {
	n := 3
	fmt.Println(generateMatrix(n))
}
