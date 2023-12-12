package main

import (
	"fmt"
)

// Brute-force approach
func maxProductBruteForce(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	mp := nums[0]

	for i := 0; i < len(nums); i++ {
		cp := 1
		for j := i; j < len(nums); j++ {
			cp *= nums[j]
			mp = max(cp, mp)
		}
	}

	return mp
}

// Optimized solution using linear traversal
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	mp, cmax, cmin := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			cmax, cmin = cmin, cmax
		}
		cmax = max(nums[i], cmax*nums[i])
		cmin = min(nums[i], cmin*nums[i])

		mp = max(cmax, mp)
	}
	return mp
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{2, 3, -2, 4}
	fmt.Println("Input:", nums1)
	fmt.Println("Optimized Output:", maxProduct(nums1))             // Optimized Output: 6
	fmt.Println("Brute-Force Output:", maxProductBruteForce(nums1)) // Brute-Force Output: 6

	nums2 := []int{-2, 0, -1}
	fmt.Println("\nInput:", nums2)
	fmt.Println("Optimized Output:", maxProduct(nums2))             // Optimized Output: 0
	fmt.Println("Brute-Force Output:", maxProductBruteForce(nums2)) // Brute-Force Output: 0
}
