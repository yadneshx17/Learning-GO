package main

import "fmt"



// this function finds an index of an element in a slice.
// [S ~[]E, E comparable]   -> generic type declaration
// means the function works with any slice type whose elements are comparable.

// comparable built-in constraint is used to specify that at type can be compared for equality. using the standarad comparision operators.
// it gets staify by all the types which supports equality operators.
// comparable is used to constrain E to types that support equality comparison.

// S ~[]E ->   means S must be the slice of E where E is the element type.~ means:

any type whose underlying type is []E
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
	next *element[T]
	val T
}

func (lst *List[T]) Push (v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val : v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val : v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	_ = SlicesIndex[[]string, string](s, "bar")

    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    fmt.Println("list:", lst.AllElements())
}
