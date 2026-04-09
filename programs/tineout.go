package main

// Goroutine = worker
// Channel = delivery of result


import (
	"fmt"
	"time"
)

// select with timeout 
// channel-based result handling


func fakeAPI() string {
	time.Sleep(2 * time.Second)
	return "data"
}

func main() {

	ch := make(chan string)

	// go routines doesnt return value, must use channel to get the result from function return.
	go func() {
		ch <- fakeAPI()
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(3 * time.Second):
		fmt.Println("timeOut")
	}

}
