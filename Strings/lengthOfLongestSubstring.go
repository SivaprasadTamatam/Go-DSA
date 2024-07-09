package main

import "fmt"

func BruteForcelengthOfLongestSubstring(s string) int {
	l := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if AllUnique(s, i, j) {
				if l < j-i+1 {
					l = j - i + 1
				}
			}
		}
	}
	return l
}

func AllUnique(s string, st, ed int) bool {
	mp := make(map[byte]bool)
	fmt.Println(string(s[st:ed]))

	for i := st; i <= ed; i++ {
		if _, ok := mp[s[i]]; ok {
			return false
		}
		mp[s[i]] = true
	}
	return true
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)

	mp := make(map[byte]int)
	res := 0
	left := 0

	for right := 0; right < n; right++ {

		if i, v := mp[s[right]]; v && i >= left {
			left = i + 1
		}
		mp[s[right]] = right
		cl := right - left + 1
		if res < cl {
			res = cl
		}
		fmt.Println(mp, left, right, res)
	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcd"))
}
