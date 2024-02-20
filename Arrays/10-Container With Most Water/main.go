package main

import (
	"fmt"
)

func maxArea(height []int) int {
	max := 0
	left := 0
	right := len(height) - 1

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if area > max {
			max = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}

func maxAreaBruteForce(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := min(height[i], height[j]) * (j - i)
			if area > max {
				max = area
			}
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxAreaBruteForce(height))
	fmt.Println(maxArea(height))
}
