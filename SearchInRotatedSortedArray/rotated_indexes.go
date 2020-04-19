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
	arr := NewRotatedSortedArray(nums)
	return arr.IndexOf(target)
}

// RotatedSortedArray stores a rotated, sorted array.
type RotatedSortedArray struct {
	values []int // Numbers
	lhs    int   // Left-hand size (smaller values)
	rhs    int   // Right-hand size (larger values)
	size   int   // Size of slice of values
}

// NewRotatedSortedArray returns a new array.
func NewRotatedSortedArray(nums []int) *RotatedSortedArray {
	rhs := findRotation(nums)
	return &RotatedSortedArray{
		values: nums,
		lhs:    len(nums) - rhs,
		rhs:    rhs,
		size:   len(nums),
	}
}

// findRotation returns ow many right rotations were performed on
// the sorted array.  It is the index of the smallest value.
func findRotation(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		med := (left + right) / 2
		if nums[med] < nums[right] {
			right = med
		} else if nums[med] > nums[right] {
			left = med + 1
		}
	}
	return left
}

// IndexOf returns the index of a value using binary search.
func (arr RotatedSortedArray) IndexOf(target int) int {
	if arr.size == 0 {
		return -1
	}
	left := 0
	right := arr.size - 1
	for left <= right {
		med := (left + right) / 2
		medVal := arr.getValue(med)
		if medVal < target {
			left = med + 1
		} else {
			right = med - 1
		}
	}
	medVal := arr.getValue(left)
	if medVal == target {
		return arr.getRotatedIndex(left)
	}
	return -1
}

func (arr RotatedSortedArray) getValue(i int) int {
	pos := arr.getRotatedIndex(i)
	return arr.values[pos]
}

func (arr RotatedSortedArray) getRotatedIndex(i int) int {
	if arr.lhs == 0 {
		return i
	}
	var pos int
	if i < arr.lhs {
		pos = i + arr.rhs
	} else {
		pos = i - arr.lhs
	}
	return pos
}
