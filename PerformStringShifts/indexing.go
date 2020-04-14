package main

import (
	"fmt"
)

func stringShift(s string, shift [][]int) string {
	n := len(s)
	var begin int
	for _, op := range shift {
		begin += (op[0]*-2 + 1) * op[1]
		for begin < 0 {
			begin = n + begin
		}
	}
	begin = begin % n
	return s[begin:] + s[:begin]
}

func main() {
	shift := [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}}
	result := stringShift("abcdefg", shift)
	fmt.Println(result, "efgabcd")
	{
		shift := [][]int{{0, 1}, {1, 2}}
		result := stringShift("abc", shift)
		fmt.Println(result, "cab")
	}
}
