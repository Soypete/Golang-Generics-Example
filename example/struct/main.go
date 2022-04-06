package main

import "fmt"

type node[T any] struct {
	Data T
	next *node[T]
	prev *node[T]
}

type list[T any] struct {
	first *node[T]
	last  *node[T]
}

func (l *list[T]) add(data T) *node[T] {
	n := &node[T]{
		Data: data,
		prev: l.last,
	}
	if l.first == nil {
		l.first = n
		l.last = n
		return n
	}
	l.last.next = n
	l.last = n
	return n
}

type user struct {
	name    string
	company string
}

func main() {
	// store users in a list
	var lv list[user]
	n1 := lv.add(user{name: "John", company: "Google"})
	n2 := lv.add(user{name: "Jane", company: "Facebook"})

	// TODO: add people from the meetup

	fmt.Println(n1.Data, n2.Data)
	fmt.Println(lv)
}
