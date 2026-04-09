// Fan-in a concurrency design pattern used to aggregate data or results from multiple independent sources into a single, unified stream
// multiple goroutine -> single stream
// Multiple independent sources ( producers ) → one output channel

// Create 2 goroutines:

// one sends numbers 1,2,3
// another sends 100,200,300
// Goal

// Merge both into one channel

package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// worker 1
	go func() {
		for _, v := range []int{1, 2, 3} {
			ch1 <- v
		}
		close(ch1)
	}()

	// worker 2
	go func() {
		for _, v := range []int{100, 200, 300} {
			ch2 <- v
		}
		close(ch2)
	}()

	// merged
	merged := make(chan int)

	go func() {
		for {
			select {
			case v, ok := <-ch1:
				if ok {
					merged <- v
				} else {
					ch1 = nil
				}
			case v, ok := <-ch2:
				if ok {
					merged <- v
				} else {
					ch2 = nil
				}
			}

			if ch1 == nil && ch2 == nil {
				close(merged)
				return
			}
		}
	}()

	// consume merged output
	for v := range merged {
		fmt.Println(v)
	}
}
