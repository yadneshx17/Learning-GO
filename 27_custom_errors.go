package main 

import (
	"errors"
	"fmt"
)

// its good to suffix custom errors with `Error` 
type argError struct {
	arg 	int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %w", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "cant work with it"}
	}
	return arg+3, nil
}

func main() {
	
}