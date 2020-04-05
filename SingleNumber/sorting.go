package main

import (
	"fmt"
	"sort"
)

func singleNumber(nums []int) int {
	sort.Ints(nums)
	var hasCandidate bool = false
	var val int
	for _, v := range nums {
		if !hasCandidate {
			hasCandidate = true
			val = v
		} else if val == v {
			hasCandidate = false
		} else {
			return val
		}
	}
	return val
}

func main() {
	nums := []int{2, 2, 1}
	n := singleNumber(nums)
	fmt.Println(n)
}
