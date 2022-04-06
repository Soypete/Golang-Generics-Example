package main

import (
	"fmt"
	"unicode/utf8"
)

type vector[T any] []T

func (v vector[T]) last() (T, error) {
	var zero T
	if len(v) == 0 {
		return zero, fmt.Errorf("empty vector")
	}
	// TODO: this probably needs to be contraint ordered
	return v[len(v)-1], nil
}

func main() {
	fmt.Println("vector[int] : ")
	vGenInt := vector[int]{10, -1}
	i, err := vGenInt.last()
	if i < 0 {
		fmt.Println("negative number")
	}
	fmt.Println("value: %d error: %v\n", i, err)

	fmt.Println("vector[string] : ")
	vGentStr := vector[string]{"hello", "world"}
	s, err := vGentStr.last()
	if !utf8.ValidString(s) {
		fmt.Println("invalid utf8 string")
	}
	fmt.Println("value: %s error: %v\n", s, err)
}
