package main

import (
	"fmt"
)

func main() {
	type Node struct {
		Next  *Node
		Value interface{}
	}

	n1 := &Node{Value: 0}
	n2 := &Node{Value: 1}
	n3 := &Node{Value: 2}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n1 // creates a cycle

	visited := map[*Node]bool{
		n1: false,
		n2: false,
		n2: false,
	}
	//
	for n := n1; n != nil; n = n.Next {
		if visited[n] { // presence of the key, can be tested like this, where non-existing key in map returns 'false' which is the zero value of this map's value type.
			fmt.Println("cycle detected")
			break
		}
		visited[n] = true
		fmt.Println(n.Value)
	}
}
