package main

import (
	"fmt"
)

func findMaxLength(nums []int) int {
	n := len(nums)
	var max int
	for i := 0; i < n; i++ {
		var zeros, ones int
		for j := i; j < n; j++ {
			if nums[j] == 0 {
				zeros++
			} else {
				ones++
			}
			if zeros == ones {
				currMax := j - i + 1
				if currMax > max {
					max = currMax
				}
			}
		}
	}
	return max
}

func main() {
	input := []int{0, 1, 0}
	result := findMaxLength(input)
	fmt.Println(result, 2)
	{
		input := []int{0, 1}
		result := findMaxLength(input)
		fmt.Println(result, 2)
	}
	{
		input := []int{0, 0, 1, 0, 0, 0, 1, 1}
		result := findMaxLength(input)
		fmt.Println(result, 6)
	}
	{
		input := []int{0, 1, 1, 0, 1, 1, 1, 0}
		result := findMaxLength(input)
		fmt.Println(result, 4)
	}
	{
		input := []int{1, 1, 1, 1, 1, 1, 1, 1}
		result := findMaxLength(input)
		fmt.Println(result, 0)
	}
}
