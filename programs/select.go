package main

import (
	"fmt"
	"time"
)

func main() {

	v1 := make(chan string, 2)
	v2 := make(chan string, 2)

	func(name string) {
		time.Sleep(2 * time.Second)
		v2 <- name
		v1 <- name
	}("yadnesh")

	select {
	case v := <-v1:
		fmt.Println("channel 1 sends", v)
	case vv := <-v2:
		fmt.Println("channel 2 sends", vv)
	default:
		fmt.Println("neither channel run")
	}

	// time.Sleep(time.Second)
}
