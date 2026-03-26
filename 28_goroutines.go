package main

import (
	"fmt"
	"time"
)

// go routines is a thread of execution
// with go the fucntions run in bg it becomes ( non-blocking )
// light weight compared to threads.
// managed to Go runtime

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// synchronous - blocking call
	f("direct")
	fmt.Println("After this output of goroutines.")

	// thsi goroutine will execute concurrent to above func call.
	go f("goroutine")

	// these two func calls are running async in separate goroutine now.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

	//When we run this program, we see the output of the blocking call first, then the output of the two goroutines.

}
