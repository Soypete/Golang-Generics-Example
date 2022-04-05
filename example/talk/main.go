package main

import (
	"fmt"
	"unicode/utf8"
)

type vector[T any] interface {
	value []T
	last() ([]T, error)
}

func (v vector[T])last[T any]() ([]T, error) {
	var zero []T
	if len(v.value) == 0 {
		return zero, fmt.Errorf("empty vector")
	}
	// TODO: this probably needs to be contraint ordered
	return v.value[len(v.value)-1], nil
}

func main() {
	fmt.Println("vector[int] : ")
	vGenInt := vector{10, -1}
	i, err := vGenInt.last()
	if i < 0 {
		fmt.Println("negative number")
	}
	fmt.Println("value: %d error: %v\n", i, err)

	fmt.Println("vector[string] : ")
	vGentStr := vector{"hello", "world"}
	s, err := vGentStr.last()
	if !utf8.ValidString(s) {
		fmt.Println("invalid utf8 string")
	}
	fmt.Println("value: %s error: %v\n", s, err)
}
