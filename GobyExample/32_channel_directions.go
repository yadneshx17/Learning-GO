package main

import (
	"fmt"
)

// channels are function parameters
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pong chan<- string) {
	msg := <-pings
	pong <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
