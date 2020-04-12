package main

import (
	"container/heap"
	"fmt"
)

type StoneHeap []int

func (h StoneHeap) Len() int {
	return len(h)
}

func (h StoneHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h StoneHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StoneHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *StoneHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := StoneHeap(stones)
	heap.Init(&h)
	for h.Len() > 1 {
		a := heap.Pop(&h).(int)
		b := heap.Pop(&h).(int)
		if a != b {
			heap.Push(&h, a-b)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(&h).(int)
}

func main() {
	input := []int{2, 7, 4, 1, 8, 1}
	result := lastStoneWeight(input)
	fmt.Println(result)
	{
		input := []int{1, 3}
		result := lastStoneWeight(input)
		fmt.Println(result)
	}
	{
		input := []int{2, 2}
		result := lastStoneWeight(input)
		fmt.Println(result)
	}
}
