package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	var res []int

	rows, cols := len(matrix), len(matrix[0])

	left, right, top, bottom := 0, cols-1, 0, rows-1

	for left <= right && top <= bottom {
		// to left

		for col := left; col <= right; col++ {
			res = append(res, matrix[top][col])
		}
		// to down

		for row := top + 1; row <= bottom; row++ {
			res = append(res, matrix[row][right])
		}

		if left < right && top < bottom {
			// to right
			for col := right - 1; col > left; col-- {
				res = append(res, matrix[bottom][col])
			}

			// to up

			for row := bottom; row > top; row-- {
				res = append(res, matrix[row][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return res
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{6, 7, 8},
	}
	fmt.Println(spiralOrder(matrix))
}
