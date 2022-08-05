package main

import (
	"fmt"
	"sync"
)

/**
--> 使用 Go 内置类型 sync.WaitGroup 来解决 channel.go 中的问题，并且优化代码，将 waitGroup.Done() 操作包装成函数
*/

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, waitGroup *sync.WaitGroup) worker {
	in := make(chan int)
	w := worker{
		in: in,
		done: func() {
			waitGroup.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func chanDemo() {
	var workers [10]worker
	var waitGroup sync.WaitGroup
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &waitGroup)
	}
	// 添加任务数量
	waitGroup.Add(20)

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// 等待任务结束
	waitGroup.Wait()
}

func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)

		// 任务结束
		w.done()
	}
}

func main() {
	chanDemo()
}
