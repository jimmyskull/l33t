package main

import (
	"errors"
	"fmt"
)

type BackspacedReader struct {
	data *string
	head int
}

func NewBackspacedReader(str *string) *BackspacedReader {
	return &BackspacedReader{
		data: str,
		head: len(*str),
	}
}

func (br *BackspacedReader) IsExhausted() bool {
	return br.head < 0
}

func (br *BackspacedReader) Head() byte {
	return (*br.data)[br.head]
}

func (br *BackspacedReader) HeadIsBackspace() bool {
	return br.head >= 0 && br.Head() == '#'
}

func (br *BackspacedReader) Next() (byte, error) {
	var counter int
	br.head--
	for br.HeadIsBackspace() || counter > 0 {
		if br.HeadIsBackspace() {
			br.head--
			counter++
		} else if counter > 0 {
			br.head--
			counter--
		}
	}
	if br.head < 0 {
		return byte(0), errors.New("exhausted")
	}
	return (*br.data)[br.head], nil
}

func backspaceCompare(S string, T string) bool {
	SReader := NewBackspacedReader(&S)
	TReader := NewBackspacedReader(&T)
	for !SReader.IsExhausted() && !TReader.IsExhausted() {
		sc, serr := SReader.Next()
		tc, terr := TReader.Next()
		if serr != nil || terr != nil {
			break
		}
		if sc != tc {
			return false
		}
	}
	return SReader.IsExhausted() && TReader.IsExhausted()
}

func main() {
	result := backspaceCompare("ab#c", "ad#c")
	fmt.Println(result)

	result = backspaceCompare("ab##", "c#d#")
	fmt.Println(result)

	result = backspaceCompare("bxj##tw", "bxo#j##tw")
	fmt.Println(result)
}
