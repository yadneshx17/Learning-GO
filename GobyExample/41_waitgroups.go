// To wait for multiple goroutines to finish, we can use a wait group.
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // simulates expensive task
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// This WaitGroup is used to wait for all the goroutines launched here to finish.
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// launch several goroutines using WaitGroup.Go
		wg.Go(func() {
			// to pass waitgroup explicitly into function, it should be by pointer.
			worker(i)
		})
	}

	// block until all the goroutines started by `wg` are dones
	wg.Wait()
	// when the function it invokes returns.
}