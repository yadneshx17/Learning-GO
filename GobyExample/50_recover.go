package main

import "fmt"

func mypani() {
	panic("a problem")
}

func main() {

	// recover() must call in deferd function
	// The return value of recover is the error raised in the call to panic.

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mypani()

	// This code will not run, because mayPanic panics. The execution of main stops at the point of the panic and resumes in the deferred closure.
	fmt.Println("After mayPanic()")
}
