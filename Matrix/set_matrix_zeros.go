package main

import "fmt"

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	rowZero := make([]bool, m)
	colZero := make([]bool, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rowZero[i] = true
				colZero[j] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rowZero[i] || colZero[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	}

	fmt.Println("Original Matrix:")
	printMatrix(matrix)

	setZeroes(matrix)

	fmt.Println("Matrix after setting zeroes:")
	printMatrix(matrix)
}
