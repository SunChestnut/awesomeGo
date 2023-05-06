package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// create a context with an expiration time of 1 second
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()

	go handler(ctx, 1500*time.Millisecond)

	select {
	// blocked
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
		//default:
		//	fmt.Println("done!!")
	}
}

func handler(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handler", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
