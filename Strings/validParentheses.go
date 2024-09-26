package main

import "fmt"

func isValid(s string) bool {

	stack := []byte{}
	mp := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {

		c, ok := mp[byte(ch)]

		if ok {
			ls := len(stack)
			if ls > 0 && stack[ls-1] == c {
				stack = stack[:ls-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, byte(ch))
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
}
