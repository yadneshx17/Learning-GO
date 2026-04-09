package main

import (
	"fmt"
	"time"
)

func main() {

	// normal blocking call
	// provides a channel
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C // blocks.
	fmt.Println("Timer 1 fired")

	// non-blocking not gets enought time
	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 Fired")
	}()

	// executes before timer2 fully executes.
	// stop2 := timer2.Stop()
	// if stop2 {
	// 	fmt.Println("Timer 2 Stopped")
	// }

	// time.Sleep(3 * time.Second)
}
