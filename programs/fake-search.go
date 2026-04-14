package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string

type Search func(query string) Result

// fake search generator
func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q", kind, query))
	}
}

// create search functions
var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplicas := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplicas(i)
	}
	return <-c
}

// Google function (concurrent search)
func Google(query string) []Result {
	c := make(chan Result)

	// Fan-in Pattern
	// launch all searches concurrently
	go func() { c <- Web(query, ) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	var results []Result

	// collect results
	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}

	return results

	// OR
	// Timeout
	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timeout")
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	results := Google("golang")

	for _, r := range results {
		fmt.Println(r)
	}
}
