package main

import (
	"fmt"
)

func main() {
	result := checkValidString("(*))")
	fmt.Println(result, true)
	result = checkValidString(")(")
	fmt.Println(result, false)
}

func checkValidString(s string) bool {
	return matchParenthesis(s, 0)
}

func matchParenthesis(s string, size int) bool {
	for i, chr := range s {
		if size < 0 {
			break
		}
		switch chr {
		case '*':
			if matchParenthesis(s[i+1:], size+1) {
				return true
			}
			if matchParenthesis(s[i+1:], size-1) {
				return true
			}
			// otherwise, treat * as null and continue to the next char.
		case '(':
			size++
		case ')':
			size--
		}
	}
	return size == 0
}
