package main

import (
	"fmt"
)

func typeString(input string) string {
	var output []rune
	for _, key := range input {
		if key == '#' {
			if len(output) > 0 {
				output = output[:len(output)-1]
			}
		} else {
			output = append(output, key)
		}
	}
	return string(output)
}

func backspaceCompare(S string, T string) bool {
	return typeString(S) == typeString(T)
}

func main() {
	result := backspaceCompare("ab#c", "ad#c")
	fmt.Println(result)
}
