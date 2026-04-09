package main

import "fmt"

func main () {
	// basic single condition
	i:=1
	for i<=3 {
		fmt.Println(i)
		i = i+1
	}

	// classic: intial/condition/after for loop.
	for j:=0; j<3; j++ {
		fmt.Println(j)
	}

	// iteration in range over an integer
	for i:= range 3 {
		fmt.Println("range", i)
		break
	}

	// for without a condition will loop repeatedly 
	// until you break out of the loop or return from the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	// can also continue to the next iteration of the loop.
	for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
