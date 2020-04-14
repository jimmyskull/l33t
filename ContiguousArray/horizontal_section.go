package main

import (
	"fmt"
)

func findMaxLength(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	y := make([]int, n)
	var maxDist int
	// Compute y = f(x), where f(x) is the difference
	// of counting 0s and 1s from the beginning.
	y[0] = nums[0]*2 - 1
	yMin, yMax := y[0], y[0]
	for i := 1; i < n; i++ {
		curr := y[i-1] + (nums[i]*2 - 1)
		y[i] = curr
		if curr > yMax {
			yMax = curr
		}
		if curr < yMin {
			yMin = curr
		}
		if curr == 0 {
			maxDist = i + 1
		}
	}
	// maxDist contains the maximum value for a subarray beginning at
	// first postition.
	//
	// Now, we seach for the largest (b-a) value where
	// f(a) = f(b), like cutting a horizontal line at any y value.
	yPos := make([]int, n)
	for i := range yPos {
		yPos[i] = -1
	}
	for i := 0; i < n; i++ {
		curr := y[i] - yMin
		firstPos := yPos[curr]
		if firstPos == -1 {
			yPos[curr] = i
		} else {
			dist := i - firstPos
			if dist > maxDist {
				maxDist = dist
			}
		}
	}
	return maxDist
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
