package main

import (
	"fmt"
)

func maxProfit(nums []int) int {
	profit := 0
	if len(nums) == 0 {
		return 0
	}
	for i, curr := range nums[1:] {
		old := nums[i]
		if curr > old {
			profit += curr - old
		}
	}
	return profit
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	n := maxProfit(nums)
	fmt.Println(n)
}
