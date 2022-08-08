package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func simpleSelectDemo() {
	c1, c2 := generator(), generator()
	// select + default 可以创建一个非阻塞式的接收方式
	for {
		select {
		case n := <-c1:
			fmt.Println(n)
		case n := <-c2:
			fmt.Println(n)
			//default:  // 从 c1 和 c2 读值的时候需要等待，因此如果加上了 default，select 就会直接选择 default 来执行
			//	time.Sleep(2 * time.Second)
			//	fmt.Println("No value received")
		}
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func worker(id int, c chan int) {
	for v := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", id, v)
	}
}

// writeAndRead 函数中 c1 和 c2 负责接收数据，而 worker 负责从 c1 或 c2 中读取数据
// 该函数的弊端在于：当 c1、c2 写入的速度大于 worker 读取的速度时候，就会导致某些数字未被读取到
func writeAndRead() {
	// c1、c2 接收数据
	c1, c2 := generator(), generator()
	// 当 c1 或 c2 中存在数据时，worker 从其中接收数据
	var worker = createWorker(0)

	n := 0
	hasValue := false
	for {
		var activeWorker chan int
		if hasValue {
			activeWorker = worker
		}
		select {
		case n = <-c1:
			hasValue = true
		case n = <-c2:
			hasValue = true
		case activeWorker <- n:
			hasValue = false
		}
	}
}

func writeAndReadBetter() {
	c1, c2 := generator(), generator()
	var worker = createWorker(1)

	var values []int
	// 从程序开始执行后开始计时，设置 10s 后停止下面操作
	tm := time.After(10 * time.Second)
	// 使用 time.Tick 设置定时执行某些操作，time.Tick 返回的是存放 Time 类型的 chan。
	// 下面示例中设置每秒钟打印一次当前队列的长度，也就是每隔 1s 往 chan 中写入当前时间
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			// 设置每两次 select 的时间间隔为 800ms 时，就打印 timeout
			fmt.Println("timeout")
		case v := <-tick:
			fmt.Printf("queue len = %d, v = %v\n", len(values), v)
		case <-tm:
			fmt.Println("bye~")
			return
		}
	}
}

func main() {
	//simpleSelectDemo()
	//writeAndRead()
	writeAndReadBetter()
}
