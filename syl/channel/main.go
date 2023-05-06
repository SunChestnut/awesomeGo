package main

import (
	"fmt"
	"time"
)

// generator 创建类型为 int 的 Channel，且向 Channel 中不断写入数字
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			//time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func receiveFromMultiChan() {
	c1, c2 := generator(), generator()
	for {
		select {
		case n := <-c1:
			fmt.Println("Received from c1:", n)
		case n := <-c2:
			fmt.Println("Received from c2:", n)
		default:
			fmt.Println("Nothing")
		}
	}
}

// produce 向 channel 中发送数据
func produce(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Printf("send...%v\n", i)
	}
	// 关闭 channel，表示发送结束
	close(ch)
}

// consume 从 channel 中接收数据
func consume(ch <-chan int) {
	//for {
	//	time.Sleep(10 * time.Millisecond)
	//	fmt.Println(<-ch)
	//}
	for v := range ch {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("receive...%v\n", v)
	}
}

func main() {

	// 创建一个 capacity=1 的 channel
	ch := make(chan int, 1)
	go produce(ch)
	consume(ch)
}
