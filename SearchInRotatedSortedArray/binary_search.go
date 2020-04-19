package main

import (
	"fmt"
)

func main() {
	var result int
	result = search([]int{5, 1, 2, 3, 4}, 1)
	fmt.Println(result, 1)
	result = search([]int{3, 5, 1}, 4)
	fmt.Println(result, -1)
	result = search([]int{5, 1, 3}, 5)
	fmt.Println(result, 0)
	result = search([]int{3, 1}, 1)
	fmt.Println(result, 1)
	result = search([]int{3, 1}, 0)
	fmt.Println(result, -1)
	result = search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	fmt.Println(result, 4)
	result = search([]int{4, 5, 6, 7, 0, 1, 2}, 3)
	fmt.Println(result, -1)
	result = search([]int{0, 1, 2, 4, 5, 6, 7}, 6)
	fmt.Println(result, 5)
	result = search([]int{}, 5)
	fmt.Println(result, -1)
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		midVal := nums[mid]
		rightVal := nums[right]
		if midVal == target {
			return mid
		}
		if midVal <= rightVal {
			if midVal <= target && target <= rightVal {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			leftVal := nums[left]
			if leftVal <= target && target <= midVal {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}
