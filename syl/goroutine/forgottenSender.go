package main

import (
	"context"
	"fmt"
	"time"
)

// The Forgotten Sender

func forgottenSender(ch chan int) {
	data := 3

	// This is a blocked as no one is receiving the data
	ch <- data
}

func handler() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelFunc()

	ch := make(chan int)
	go forgottenSender(ch)

	go func() {
		for {
			select {
			case data := <-ch:
				fmt.Println(data)
			case <-ctx.Done():
				fmt.Println("Timeout! Processor cancelled. Returning")
				return // using `return` instead of `break` to exit the entire `for` loop
			}
		}
	}()
}

func main() {
	handler()

	time.Sleep(10 * time.Second)
	fmt.Println("main exist.")
}
