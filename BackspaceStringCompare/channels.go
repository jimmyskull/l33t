package main

import (
	"fmt"
)

func reverseString(c chan byte, data *string) {
	head := len(*data)
	var counter int
	for head > 0 {
		head--
		if (*data)[head] == '#' {
			counter++
		} else if counter > 0 {
			counter--
		} else {
			c <- (*data)[head]
		}
	}
	close(c)
}

func backspaceCompare(S string, T string) bool {
	s, t := make(chan byte), make(chan byte)

	go reverseString(s, &S)
	go reverseString(t, &T)

	var sVal, tVal byte
	var sDone, tDone bool
	for !sDone && !tDone {
		sVal = <-s
		tVal = <-t
		if sVal != tVal {
			return false
		}
		sDone, tDone = sVal == byte(0), tVal == byte(0)
	}
	return sDone && tDone
}

func main() {
	result := backspaceCompare("ab#c", "ad#c")
	fmt.Println(result, true)

	result = backspaceCompare("ab##", "c#d#")
	fmt.Println(result, true)

	result = backspaceCompare("bxj##tw", "bxo#j##tw")
	fmt.Println(result, true)

	result = backspaceCompare("bxj##tw", "bxo#j##w")
	fmt.Println(result, false)

	result = backspaceCompare("bxj##tw", "bxj###tw")
	fmt.Println(result, false)
}
