// rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service.
package main

import (
	"fmt"
	"time"
)

func main() {

	// requests := make(chan int, 5)

	// for i := 1; i <=5; i++ {
	// 	requests <- i
	// }
	// close(requests)

	// // this limiter channel will receive a value every 200 milliseconds.
	// limiter := time.Tick(200 * time.Millisecond)

	// for req := range requests {
	// 	// By blocking on a receive from the limiter channel before serving
	// 	// each request, we limit ourselves to 1 request every 200 milliseconds.
	// 	<- limiter
	// 	fmt.Println("request", req, time.Now())
	// }

	// This burstyLimiter channel will allow bursts of up to 3 events.
	burstyLimiter := make(chan time.Time, 3)

	for range 3 {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
