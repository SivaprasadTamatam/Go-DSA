package main

import (
	"fmt"
	"sort"
)

func threeSumTwoPointer(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			total := nums[i] + nums[left] + nums[right]

			if total < 0 {
				left++
			} else if total > 0 {
				right--
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}

	return result
}

func threeSumBruteForce(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	length := len(nums)

	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			for k := j + 1; k < length; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					triplet := []int{nums[i], nums[j], nums[k]}
					if !contains(result, triplet) {
						result = append(result, triplet)
					}
				}
			}
		}
	}

	return result
}

func contains(result [][]int, triplet []int) bool {
	for _, r := range result {
		if equal(r, triplet) {
			return true
		}
	}
	return false
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println("Using Two-pointer Technique Solution:")
	fmt.Println(threeSumTwoPointer(nums)) // Output: [[-1 -1 2] [-1 0 1]]

	fmt.Println("\nUsing Brute Force Solution:")
	fmt.Println(threeSumBruteForce(nums)) // Output: [[-1 -1 2] [-1 0 1]]
}
