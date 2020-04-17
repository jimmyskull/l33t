package main

import (
	"fmt"
)

func main() {
	{
		grid := [][]byte{}
		result := numIslands(grid)
		fmt.Println(result, 0)
	}
	{
		grid := [][]byte{
			{'1','1','1','1','0'},
			{'1','1','0','1','0'},
			{'1','1','0','0','0'},
			{'0','0','0','0','0'},
		}
		result := numIslands(grid)
		fmt.Println(result, 1)
	}
	{
		grid := [][]byte{
			{'1','1','0','0','0'},
			{'1','1','0','0','0'},
			{'0','0','1','0','0'},
			{'0','0','0','1','1'},
		}
		result := numIslands(grid)
		fmt.Println(result, 3)
	}
}

func numIslands(grid [][]byte) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	var islands int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visitIsland(grid, i, j) {
				islands++
			}
		}
	}
    return islands
}

func visitIsland(grid [][]byte, i, j int) bool {
	if getGridValue(grid, i, j) != '1' {
		return false
	}
	grid[i][j] = '0'
	visitIsland(grid, i-1, j)
	visitIsland(grid, i+1, j)
	visitIsland(grid, i, j-1)
	visitIsland(grid, i, j+1)
	return true
}

func getGridValue(grid [][]byte, i, j int) byte {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return '0'
	}
	return grid[i][j]
}
