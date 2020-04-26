package main

import "fmt"

func main() {
	{
		r := canJump([]int{2, 3, 1, 1, 4})
		fmt.Println(r, true)
	}
	{
		r := canJump([]int{3, 2, 1, 0, 4})
		fmt.Println(r, false)
	}
}

func canJump(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return true
	}
	max := nums[0]
	for i := 0; i < n; i++ {
		if max <= i && nums[i] == 0 {
			return false
		}
		max = maxVal(max, i+nums[i])
		if max >= n-1 {
			return true
		}
	}
	return false
}

func maxVal(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
