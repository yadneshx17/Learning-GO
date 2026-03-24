package main

import (
	"fmt"
	"math"
)

// interfaces -> named collection of method signatures.

// interface for geometry shapes
type geometry interface {
	area() float64
	peri() float64
}

type rectt struct {
	width, height float64
}

type circle struct {
	radius float64
}

// rect
func (r rectt) area() float64 {
	return r.width * r.height
}
func (r rectt) peri() float64 {
	return 2*r.width + 2*r.height
}

// circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) peri() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.peri())
}

// type assertions
// g.(circle) -> means extrect the underlying value from interface g as a circle.
// do something -> store value -> check condition
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok { // c -> the actual value. ok -> boolean(true/false)
		fmt.Println("Circle with radius", c.radius)
	}
}

func main() {
	r := rectt{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
