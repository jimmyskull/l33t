package main

import (
	"fmt"

	"github.com/golang-collections/go-datastructures/bitarray"
)

func singleNumber(nums []int) int {
	exist := bitarray.NewSparseBitArray()
	dup := bitarray.NewSparseBitArray()

	for _, v := range nums {
		val := uint64(v)
		valExist, err := exist.GetBit(val)
		if err != nil {
			panic(err)
		}
		if valExist {
			dup.ClearBit(val)
		} else {
			exist.SetBit(val)
			dup.SetBit(val)
		}
	}
	return int(exist.And(dup).ToNums()[0])
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	n := singleNumber(nums)
	fmt.Println(n)
}
