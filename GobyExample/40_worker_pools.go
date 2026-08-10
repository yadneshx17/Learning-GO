package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Why jobs are filled after Spawing up the workers?
	// A send to a buffered channel blocks when the buffer is full.
	// so if 10 jobs buffer cap is 5
	// jobs <- 6  at this buffered channel is full (5), worker yet to be spawned to finish up the jobs
	// main goroutine is stucked at jobs <- 6
	// it never reaches worker spawn logic
	// Gets a Deadlock.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		// <-results
		val := <-results
		fmt.Println("Received: ", val)
	}
}
