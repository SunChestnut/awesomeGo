package main

import (
	"fmt"
	"time"
)

// worker 函数中将 channel 作为参数
func worker(id int, c chan int) {
	for {
		fmt.Printf("Worker %d received %c\n", id, <-c)
	}
}

func simpleChanDemo() {
	// 创建类型为 channel 且 channel 中的内容为 int 类型的变量 c，但 c 的值为 nil
	//var c chan int

	// 创建可以直接用的 channel
	c := make(chan int, 2)

	// 向 channel 中发送数据
	c <- 1
	c <- 2

	// 从 channel 中接收数据
	n := <-c
	fmt.Println(n)

	go worker(0, c)

	time.Sleep(time.Millisecond)
}

// 创建只有【接收数据】功能的 channel，也就是说，接收该 channel 只能给它发数据，相对应的，在函数体内部就只能收数据
func createSendWorker(id int) chan<- int {
	// 创建 channel
	c := make(chan int)
	// 创建接收 channel 的 goroutine
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()

	return c
}

func useSendChannel() {
	// 将 channel 作为数组的类型; chan<- 表明该 channel 只能用于【发数据】
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createSendWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

// 创建只有【发送数据】功能的 channel
func createReceiveWorker() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 887; i++ {
			c <- i
		}
	}()
	return c
}

func useReceiveWorker() {
	c := createReceiveWorker()
	for {
		fmt.Println(<-c)
	}
}

// bufferedChannel 创建缓冲容量为 3 的channel，功能为：当前没有接收者的话也不会发生 deadlock
func bufferedChannel() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
}

// chanelClose 如果数据有明确结尾的话，发送方可以设置发完数据后关闭 channel
// ⚠️ 只有发送方 close 之后，接收方才可与选择接收方式，接收方不能主动 close
func channelClose() {
	// 因为有接收者，所以设不设置缓存均可
	c := make(chan int)
	go workerWithClose(0, c)
	c <- 111
	c <- 222
	c <- 333
	close(c)

	time.Sleep(time.Millisecond)
}

func workerWithClose(id int, c chan int) {
	// 如果发送方设置了发送完数据后关闭 channel，但接收方无法立即停止接受，接收方仍然会接收 1ms 时间的数据
	// ==> 如果 channel 存储的值为 int 类型，则接收者会继续收到 0，如果为 string 类型，则会接收空串
	// ==> 可以通过一下两种方式设置当发送方停止发送后，接收方也停止接收

	// 方式一：直接判断从 channel 中读出来的是否还有值，n 为读出来的值的内容，ok 为是否还有值
	//for {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d received %d\n", id, n)
	//}

	// 方式二：直接遍历 c，等到 c 的内容发完，便跳出循环
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func main() {
	fmt.Println("Channel as first-class citizen")
	simpleChanDemo()
	useSendChannel()
	useReceiveWorker()

	fmt.Println("Buffered channel")
	bufferedChannel()

	fmt.Println("Channel close and range")
	channelClose()
}
