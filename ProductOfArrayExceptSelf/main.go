package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	n := len(nums)
	left := make([]int, n+1)
	right := make([]int, n+1)
	left[0] = 1
	right[n] = 1
	for i := 0; i < n; i++ {
		left[i+1] = left[i] * nums[i]
	}
	for j := n - 1; j >= 0; j-- {
		right[j] = right[j+1] * nums[j]
	}
	output := make([]int, n)
	for i := range output {
		output[i] = left[i] * right[i+1]
	}
	return output
}

func main() {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	fmt.Println(result, []int{24, 12, 8, 6})
}
