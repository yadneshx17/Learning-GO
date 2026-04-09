package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

type List[T] struct {
	head, tail *List[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

// singly linked List
func (lst *List[T]) Push(val T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: val}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: val}
		lst.tail = lst.tail.next
	}
}

