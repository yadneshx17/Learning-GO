package main

import (
	"fmt"
	"slices"
)

func main() {

	// uninitialized slice equals to nil and has length 0
	var s []string
	fmt.Println("uninit: ", s, s == nil, len(s) == 0) // uinit:  [] true true

	// builtin make fuction to create a non-zero length slice, zero valued
	s = make([]string, 3)
	fmt.Println("emp:", s, "len", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	// slice operator: slice[low:high]
	l := s[:2]
	fmt.Println("sl1: ", l)

	l = s[2:5]
	fmt.Println("sl3", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// declare and initialize a variable for slice in single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j 
		}
	}
	fmt.Println("2d: ", twoD)

}
