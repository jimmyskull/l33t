package main

import (
	"fmt"
	"math"
)

func isHappy(n int) bool {
	var digit int
	var newNum int
	if n == 1 {
		return true
	} else if n == 4 {
		return false
	}
	for n != 0 {
		digit = int(math.Mod(float64(n), 10.0))
		n = n / 10
		newNum += digit * digit
	}
	return isHappy(newNum)
}

func main() {
	r := isHappy(19)
	fmt.Println(r)
}
