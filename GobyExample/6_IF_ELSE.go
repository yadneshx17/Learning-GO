package main

import "fmt"

func main () {
	
	// basic 
	if 7%2 == 0 {
		fmt.Println("7 is Even")
	} else {
		fmt.Println("7 is odd")
	}
	
	// only if
	if 8%4==0 {
		fmt.Println("8 is divisible by 4")
	}
	
	// logical operators
	if 8%2==0 || 7%2==0 {
		fmt.Println("Either 8 or 7 are even")
	}

	// A statement can precede conditionals; any variables declared in this statement 
	// are available in the current and all subsequent branches.
	if num := 9; num < 0 {
		fmt.Println(num, "is Negative")	
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}