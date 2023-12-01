package main

import "fmt"

// O(N^2) Time Complexity Solution
func brute_force_twoSum(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

// O(N) Time & O(N) Space Complexity Solution
func twoSum(nums []int, target int) []int {
	l := len(nums)

	mp := make(map[int]int)
	for i := 0; i < l; i++ {
		diff := target - nums[i]
		if j, ok := mp[diff]; ok {
			return []int{j, i}
		}
		mp[nums[i]] = i
	}
	return []int{-1, -1}
}

func main() {
	// Example 1
	fmt.Println(brute_force_twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

	//Example 2:
	fmt.Println(brute_force_twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))

	//Example 3:
	fmt.Println(brute_force_twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
