package main

import (
	"fmt"
)

func countElements(arr []int) int {
	table := make(map[int]bool)
	for _, num := range arr {
		table[num] = true
	}
	var count int
	for _, num := range arr {
		if _, ok := table[num+1]; ok {
			count++
		}
	}
	return count
}

func main() {
	values := []int{1, 2, 3}
	result := countElements(values)
	fmt.Println(result)
}
