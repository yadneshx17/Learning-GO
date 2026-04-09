package main

import (
	"fmt"
	"time"
)

// goroutines
// channels
// buffering
// worker pool

// worker listens → pulls jobs → processes continuously
func worker(id int, jobs <-chan int, results chan<- int) {
	// here this loop keeps reading for incoming jobs, only stops when channel is closed.
	for j := range jobs {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Worker", id, "processing", j)
		results <- j * j
	}
}

func main() {
	// invoke three workers
	const numJobs = 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// invoked the worker three times
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// for the three workers give the 10 jobs independently to the workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // closes jobs channel

	for a := 1; a <= numJobs; a++ {
		// <-results
		val := <-results
		fmt.Println(val)
	}
}
