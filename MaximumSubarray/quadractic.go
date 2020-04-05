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
	var val int
	n := len(nums)
	mat := make([]int, n)
	max := MinInt
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				val = nums[i]
				mat[j] = val
				if val > max {
					max = val
				}
			} else if j > i {
				a, b := MinInt, MinInt
				if j > 0 {
					b = mat[j-1] + nums[j]
				}
				val = int(math.Max(float64(a), float64(b)))
				mat[j] = val
				if val > max {
					max = val
				}
			}
		}
	}
	return max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	n := maxSubArray(nums)
	fmt.Println(n)
}
