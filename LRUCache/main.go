package main

import (
	"fmt"
)

func main() {
	obj := Constructor(2)
	p := obj.Get(1)
	fmt.Println(p, -1)
	obj.Put(1, 2)
	p = obj.Get(1)
	fmt.Println(p, 2)
}

type LRUCache struct {
	table    map[int]int
	order    []int
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		table:    make(map[int]int),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	value, ok := this.table[key]
	if !ok {
		return -1
	}
	this.removeFromOrder(key)
	this.addToOrder(key)
	return value
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.table[key]
	if ok {
		this.remove(key)
	}

	if this.isFull() {
		this.remove(this.order[0])
	}

	this.table[key] = value
	this.addToOrder(key)
}

func (this LRUCache) isFull() bool {
	return this.capacity == len(this.table)
}

func (this *LRUCache) remove(key int) {
	this.removeFromTable(key)
	this.removeFromOrder(key)
}

func (this *LRUCache) removeFromTable(key int) {
	delete(this.table, key)
}

func (this *LRUCache) removeFromOrder(key int) {
	n := len(this.order)
	for i := 0; i < n; i++ {
		if this.order[i] == key {
			this.order = remove(this.order, i)
			break
		}
	}
}

func (this *LRUCache) addToOrder(key int) {
	this.order = append(this.order, key)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
