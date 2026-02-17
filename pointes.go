package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

// has reference to the memory addr of the var.
func zeroptr(iptr *int) { // takes a int pointer
	// changing the value of derefrenced pointer changes the value at the referenced address
	*iptr = 0 // dereference the pointer from its memory address to the current value at that address.
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	// arguments are passed by value
	// copy of i is passed
	// distinct from the one in the calling function
	zeroval(i)
	fmt.Println("zeroval:", i)

	// gives the memory address of i, i.e a pointer to i.
	zeroptr(&i)
	fmt.Println("zeroptr:", i) // prints the memory address

	fmt.Println("pointer:", &i)
}
