package main

import (
	"fmt"
	"maps"
	// "maps"
)

func main() {

	// make(map[key-type]val-type).
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("maps:", m) // maps: map[k1:7 k2:13]

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("Len:", len(m)) // returns number of key-value pair
	
	delete(m, "k2")
	fmt.Println("maps:", m) 
	
	clear(m)
	fmt.Println("map:", m)
	
	_, prs := m["k2"] // blank idenitfier
	fmt.Println("prs:", prs)
	
	n := map[string]int{"foo":1, "bar":24}
	fmt.Println("map:", n)
	
	n2 := map[string]int{"foo":1, "bar":24}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
	
	
}
