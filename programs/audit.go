package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	alice = 1000
	bob   = 1000
	total = alice + bob

	mu sync.Mutex
)

func main() {

	// Alice -> Bob transfers
	go func() {
		for i := 0; i < 1000000; i++ {
		// atomic 

			mu.Lock()
			alice -= 1
			// mu.Unlock()

			// mu.Lock()
			bob += 1
			mu.Unlock()
		}
	}()

	// Bob -> Alice transfers
	go func() {
		for i := 0; i < 1000000; i++ {
			mu.Lock()
			bob -= 1
			// mu.Unlock()

			// mu.Lock()
			alice += 1
			mu.Unlock()
		}
	}()

	// Audit thread
	start := time.Now()

	for time.Since(start) < 1*time.Second {
		mu.Lock()

		if alice+bob != total {
			fmt.Printf(
				"Violation! alice=%d bob=%d sum=%d\n",
				alice,
				bob,
				alice+bob,
			)
		}

		mu.Unlock()
	}

	fmt.Println("Audit complete")
}