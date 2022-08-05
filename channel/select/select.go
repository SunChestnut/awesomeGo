package main

import (
	"fmt"
)

func generator() {
	out := make(chan int)
	go func() {
		for {
			i := 0
			out <- i
			i++
		}
	}()
}

func simpleSelectDemo() {
	var c1, c2 chan int
	// select + default 可以创建一个非阻塞式的接收方式
	select {
	case n := <-c1:
		fmt.Println("Received from c1: ", n)
	case n := <-c2:
		fmt.Println("Received from c2: ", n)
	default:
		fmt.Println("No value received")
	}
}

func main() {
	simpleSelectDemo()
}
