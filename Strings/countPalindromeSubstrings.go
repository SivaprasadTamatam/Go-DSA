package main

import "fmt"

func IsPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}
func countSubstringsBruteForce(s string) int {
	// O(n^3)
	res := 0
	n := len(s)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if IsPalindrome(s[i : j+1]) { // O(N)
				res++
			}
		}
	}
	return res
}

func countSubstringsOptimized(s string) int {
	res := 0
	n := len(s)

	// Expand around each center
	for c := 0; c < 2*n-1; c++ {
		l := c / 2
		r := l + c%2

		// Expand while the characters match
		for l >= 0 && r < n && s[l] == s[r] {
			res++
			l--
			r++
		}
	}

	return res
}

func main() {
	fmt.Println("Brute Force Result:", countSubstringsBruteForce("abc"))
	fmt.Println("Optimized Result:", countSubstringsOptimized("aaa"))
}
