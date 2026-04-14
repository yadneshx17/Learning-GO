package main

// An atomic counter is a variable that can be safely updated by multiple goroutines without locks

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	// counter state
	var counter atomic.Uint64

	var wg sync.WaitGroup

	start := time.Now()
	for range 50 { // 50 goroutines
		wg.Go(func() {
			for range 10000 {
				counter.Add(1)
			}
		})
	}

	wg.Wait()

	fmt.Println("ops:", ops.Load())
	elasped := time.Since(start)
	fmt.Println(elasped.Milliseconds(), "ms")
}
