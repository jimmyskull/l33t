package main

import (
	"container/list"
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

// Pair holds a (key, value) pair.
type Pair struct {
	key   int
	value int
}

// LRUCache is a Least Recently Usage (LRU) Cache context.
type LRUCache struct {
	table    map[int]*list.Element
	data     *list.List
	capacity int
}

// Constructor returns a new LRUCache.
func Constructor(capacity int) LRUCache {
	return LRUCache{
		table:    make(map[int]*list.Element, capacity),
		data:     list.New(),
		capacity: capacity,
	}
}

// Get returns the value of a key in cache, -1 if key not present.
func (l *LRUCache) Get(key int) int {
	elem, ok := l.table[key]
	if !ok {
		return -1
	}
	l.data.MoveToFront(elem)
	return valueOfElement(elem)
}

// Put maps key to value in cache, removing the least used key when the
// cache is at full capacity.
func (l *LRUCache) Put(key int, value int) {
	elem, ok := l.table[key]
	if !ok {
		if l.isFull() {
			elem = l.data.Back()
			delete(l.table, elem.Value.(Pair).key)
		} else {
			elem = l.data.PushFront(Pair{})
		}
	}
	l.data.MoveToFront(elem)
	elem.Value = Pair{key: key, value: value}
	l.table[key] = elem
}

// isFull returns whether the cache is at full capacity.
func (l LRUCache) isFull() bool {
	return l.data.Len() == l.capacity
}

func valueOfElement(e *list.Element) int {
	return pairFromElement(e).value
}

func pairFromElement(e *list.Element) Pair {
	return e.Value.(Pair)
}
