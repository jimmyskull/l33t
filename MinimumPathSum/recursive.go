package main

import (
	"fmt"
)

func main() {
	{
		grid := [][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		}
		result := minPathSum(grid)
		fmt.Println(result, 7)
	}
	{
		grid := [][]int{
			{1,2,5},
			{3,2,1},
		}
		result := minPathSum(grid)
		fmt.Println(result, 6)
	}
}

func minPathSum(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}
	mem := make([][]int, n)	
	for i := range mem {
		mem[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			right := getValue(mem, i, j-1)
			bottom := getValue(mem, i-1, j)
			mem[i][j] = minValue(right, bottom) + grid[i][j]
		}
	}
    return mem[n-1][m-1]
}

func getValue(mem [][]int, i, j int) int {
	if i < 0 || j < 0 {
		return -1
	}
	return mem[i][j]
}

func minValue(a int, b int) int {
	aMissing := a == -1
	bMissing := b == -1
	if aMissing && bMissing {
		return 0
	}
	if aMissing && !bMissing {
		return b
	}
	if !aMissing && bMissing {
		return a
	}
	if a < b {
		return a
	}
	return b
}