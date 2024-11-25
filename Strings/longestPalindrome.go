/*
Given a string s, return the longest
palindromic

substring
 in s.



Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"


Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters.
*/

package main

import "fmt"

func isPalindrome(s string, st, ed int) bool {
	for st < ed {
		if s[st] != s[ed] {
			return false
		}
		st++
		ed--
	}
	return true
}
func BruteForcelongestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	res := ""
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// Check if the substring s[i:j+1] is a palindrome
			if isPalindrome(s, i, j) {
				if j-i+1 > len(res) {
					res = s[i : j+1]
				}
			}
		}
	}
	return res
}

func OptimizedLongestPalindrome(s string) string {
	ml, st, n := 1, 0, len(s)

	if n == 1 {
		return s
	}

	l, r := 0, 0
	for i := 1; i < n; i++ {
		l = i - 1
		r = i

		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 < ml {
				ml = r - l + 1
				st = l
			}
			l--
			r++
		}

		l = i - 1
		r = i + 1

		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 < ml {
				ml = r - l + 1
				st = l
			}
			l--
			r++
		}
	}
	return s[st : st+ml]
}

func main() {
	// Test cases
	fmt.Println(OptimizedLongestPalindrome("babad")) // Output: "bab" or "aba"
	fmt.Println(BruteForcelongestPalindrome("cbbd")) // Output: "bb"
	fmt.Println(OptimizedLongestPalindrome("a"))     // Output: "a"
	fmt.Println(BruteForcelongestPalindrome("ac"))   // Output: "a" or "c"
}
