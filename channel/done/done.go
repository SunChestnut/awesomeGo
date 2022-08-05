package main

import "fmt"

/**
--> 解决 channel.go 中的问题：发送方发送数据后，因为不确定接收方何时能接收完数据，因此需要在函数中设置休眠来等待接收方接收数据
--> 解决方案：接收方接收完数据后，通知发送方
*/

type worker struct {
	in  chan int
	out chan bool
}

func createWorker(id int) worker {
	in := make(chan int)
	out := make(chan bool)
	go doWorker(id, in, out)
	return worker{in: in, out: out}
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
		// 不在这里接收接收方的反馈的原因：会变成顺序访问，也就是发送方发送一个，接收方回复一次，然后发送方再发送下一个，这样 goroutine 的创建就么的意义了
		//<-worker.out
	}

	// channel 的发送为阻塞式，任务发送完后必须被接收，才能继续后面的发送
	for _, worker := range workers {
		<-worker.out
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
		//<-worker.out
	}

	for _, worker := range workers {
		<-worker.out
	}

}

func doWorker(id int, in chan int, out chan bool) {
	for n := range in {
		fmt.Printf("Worker %d received %c\n", id, n)
		out <- true
	}
}

func main() {
	chanDemo()
}
