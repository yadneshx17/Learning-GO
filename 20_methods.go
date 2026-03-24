package main

import "fmt"

// Methods - are functions that are associated with a specific type, known as the receiver.

// You may want to use a pointer receiver type to avoid copying on method calls \
// or
// to allow the method to mutate the receiving struct.

type rect struct {
	width, height int
}

func (r *rect) area() int { // receiver type -> pointer
	return r.width * r.height
}

func (r rect) perimeter() int { // receiver type -> value
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perimeter: ", r.perimeter())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perimeter: ", rp.perimeter())
}
