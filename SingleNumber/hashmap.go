package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	m := make(map[int]int, len(nums))
	u := make(map[int]bool, len(nums))
	var curr int
	for _, v := range nums {
		curr = m[v]
		m[v] = curr + 1
		if curr == 0 {
			u[v] = true
		} else {
			delete(u, v)
		}
	}
	for k := range u {
		return k
	}
	return -1
}

func main() {
	nums := []int{2, 2, 1}
	n := singleNumber(nums)
	fmt.Println(n)
}
