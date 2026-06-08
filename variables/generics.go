package main

import "fmt"

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next  *element[T]
	value T
}

func (lst *List[T]) Push(v T) {
	// base case
	if lst.tail == nil {
		lst.head = &element[T]{value: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{value: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) Pop() (T, bool) {
	if lst.head == nil {
		var zero T
		return zero, false
	}
	value := lst.head.value
	lst.head = lst.head.next
	if lst.head == nil {
		lst.tail = nil
	}
	return value, true
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.value)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	fmt.Println("index of zoo: ", SlicesIndex(s, "zoo"))

	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(23)
	lst.Push(44)

	fmt.Println("list elements: ", lst.AllElements())

}
