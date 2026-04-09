package main

// https://go.dev/blog/pipelines

import (
	"fmt"
	"sync"
)

// func gen(nums ...int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for _, n := range nums {
// 			out <- n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

func gen(nums ...int) <-chan int {
	out := make(chan int, 2)
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done: // listen for cancellation
				return
			}
		}
		// close(out) 
	}()
	return out
}

// combine many streams  -> into one stream
// read-only channels
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	// wg counter = 3

	//	Done → 2
	//	Done → 1
	//	Done → 0

	//	→ Wait() unblocks
	//	→ close(out)

	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		defer wg.Done() // defer = run this when function exits	
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
		wg.Done() // counter -= 1 of the goroutine
	}
	// wg.Add(len(cs)) // 2 goroutines.

	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait() // blocks until counter become 0. waits for all the gorourinte
		close(out)
	}()

	return out
}

func main() {
	// Setup the pipeline.
	// c := gen(2, 3) // channel that emits the integer in the list.
	// out := sq(c)

	// // Consume the output.
	// fmt.Println(<-out) // 4
	// fmt.Println(<-out) // 6

	// for n := range sq(sq(gen(2, 3))) {
	// 	fmt.Println(n) // 16 then 81
	// }

	// Fan-Out, Fan-In
	// Multiple functions can read from the same channel until
	// that channel is closed; this is called fan-out. This provides a way to
	// distribute work amongst a group of workers to parallelize CPU use and
	// I/O.

	// A function can read from multiple inputs and proceed until all are
	// closed by multiplexing the input channels onto a single channel that’s
	// closed when all the inputs are closed. This is called fan-in.

	in := gen(2, 3)
	done := make(chan struct{}, 2)

	// fan-out
	ch1 := sq(done, in)
	// ch2 := sq(done, in)

	// Consume the first value from output
	out := merge(done, ch1, ch2)
	fmt.Println(<-out)

	// tell the remaining senders we're leaving
	// done <- strcut{}
	// done <- strcut{}
    
	// fan-in
	// Consume the merged output from c1 and c2.
	// range on channels receives only one value per iteration.
	for n := range merge(done, ch1, ch2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}
