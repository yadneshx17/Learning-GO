package main 

import (
	"errors"
	"fmt"
)

// error.New -> contstructs a basic error value with the given error message.

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}
	return arg + 3, nil // nil indicates that their is no error value
}

// A sentinel error is a predeclared variable that is used to signify a specific error condition.
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("Can't boil water")

func makeTea(arg int) error { 
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i :=range []int{7, 42} {
		if r, e:= f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}
	
	for i:= range 5 {
		if err:= makeTea(i); err!= nil {
			
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now if is dark.")
			} else {
				fmt.Println("unknown error: %s\n", err)
			}
			continue
		}
		fmt.Println("Tea is ready!")
	}
}