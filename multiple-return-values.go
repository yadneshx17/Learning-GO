package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// take specific subset of the returned values, use the blank identifier _.
	_, c := vals()
	fmt.Println(c)
}