package main

import (
	"fmt"
	"math"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func maxSubArray(nums []int) int {
	globalMax := MinInt
	localSum := 0
	for _, val := range nums {
		if localSum < 0 {
			localSum = 0
		}
		localSum = localSum + val
		if localSum > globalMax {
			globalMax = localSum
		}
	}
	return globalMax
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	n := maxSubArray(nums)
	fmt.Println(n)
	{
		nums := []int{-1}
		n := maxSubArray(nums)
		fmt.Println(n)
	}
	{
		nums := []int{-2, 1}
		n := maxSubArray(nums)
		fmt.Println(n)
	}
}
