package main

import (
	"fmt"
	"math"
)

// O(N^2) approach
func maxProfitBruteForce(prices []int) int {
	mp := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			cp := prices[j] - prices[i]
			if cp > mp {
				mp = cp
			}
		}
	}
	return mp
}

// O(N) approach
func maxProfit(prices []int) int {

	minPrice := math.MaxInt32
	mp := 0
	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		} else if (p - minPrice) > mp {
			mp = p - minPrice
		}
	}
	return mp
}

func main() {

	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Brute Force Approach")
	fmt.Println("Input:", prices)
	fmt.Println("Output:", maxProfitBruteForce(prices)) // Output: 5

	fmt.Println("Optimized Approach")
	fmt.Println("Input:", prices)
	fmt.Println("Output:", maxProfit(prices)) // Output: 5
}
