package main

import "fmt"

// channels are pipes that let two goroutines talk each other.
// use to sync go routines

func sender(ch chan<- int) {
	ch <- 100
}

func receiver(ch <-chan int) {
	val := <-ch
	fmt.Println("Received: ", val)
}

func main() {

	// creates a channel
	// channel_name := make(chan typevalue)
	messages := make(chan string)
	ch := make(chan int)

	go sender(ch)
	go receiver(ch)

	// send a value into a channel.
	go func() { messages <- "ping" }()

	// receives a value from channel
	// msg gets value from messages channel
	// which sent by the anon function.
	msg := <-messages
	fmt.Println(msg)

	var input string
	fmt.Scanln(&input)
}
