package main

import (
	"fmt"
)

// Brute Force Approach (O(N^2) time complexity
func bruteForceProductExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		product := 1
		for j := 0; j < n; j++ {
			if i != j {
				product *= nums[j]
			}
		}
		result[i] = product
	}

	return result
}

// Optimized Solution O(N) space and O(N) time
func optimizedProductExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	leftProduct := make([]int, n)
	rightProduct := make([]int, n)

	leftProduct[0] = 1
	rightProduct[n-1] = 1

	for i := 1; i < n; i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1]
	}

	for i := n - 2; i >= 0; i-- {
		rightProduct[i] = rightProduct[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		result[i] = leftProduct[i] * rightProduct[i]
	}

	return result
}

// Constant Space Approach without Division O(1) Space and O(N) time
func constantSpaceProductExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	product := 1
	for i := 0; i < n; i++ {
		result[i] = product
		product *= nums[i]
	}

	product = 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= product
		product *= nums[i]
	}

	return result
}

// Zero Count and Zero Index Solution O(N) time complexity
func zeroCountProductExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	prod := 1
	zeroCount := 0
	zeroIndex := 0

	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			prod *= nums[i]
		} else {
			zeroCount++
			zeroIndex = i
		}
	}

	if zeroCount > 1 {
		return res
	} else if zeroCount == 1 {
		res[zeroIndex] = prod
		return res
	}

	for i := 0; i < n; i++ {
		res[i] = prod / nums[i]
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4}

	fmt.Println("Brute Force Approach:", bruteForceProductExceptSelf(nums))
	fmt.Println("Optimized Solution:", optimizedProductExceptSelf(nums))
	fmt.Println("Constant Space Approach without Division:", constantSpaceProductExceptSelf(nums))
	fmt.Println("Zero Count and Zero Index Solution:", zeroCountProductExceptSelf(nums))
}
