package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker Stop or Stop the task its kinda timeout or something")
			return

		default:
			fmt.Println("Processing Task")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // releases resources associated with context task

	// go worker(ctx)

	mySleepandTalk(ctx, 5*time.Second, "Hello")
}

func mySleepandTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
