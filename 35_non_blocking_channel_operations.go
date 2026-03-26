package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan string)

	// non-blocking receive if a value is available then it will receive or default if not
	select {
	case msg := <-messages:
		fmt.Println("received Message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hello"

	// non-blocking send.
	// msg cannot be send to messages channel, because the channel has no buffer amd there is no receiver. default is selected.
	select {
	case messages <- msg:
		fmt.Println("sent messages", msg)
	default:
		fmt.Println("no messeage semmt")
	}

	// multiple cases above the default clause to implement a multi-way non-blocking select.
	select {
	case msg := <-messages:

		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
