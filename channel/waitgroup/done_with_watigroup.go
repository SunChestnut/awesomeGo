package main

import (
	"fmt"
	"sync"
)

/**
--> 使用 Go 内置类型 sync.WaitGroup 来解决 channel.go 中的问题
*/

type worker struct {
	in   chan int
	done *sync.WaitGroup
}

func createWorker(id int, waitGroup *sync.WaitGroup) worker {
	in := make(chan int)
	go doWorker(id, in, waitGroup)
	return worker{in: in, done: waitGroup}
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

func doWorker(id int, in chan int, waitGroup *sync.WaitGroup) {
	for n := range in {
		fmt.Printf("Worker %d received %c\n", id, n)

		// 任务结束
		waitGroup.Done()
	}
}

func main() {
	chanDemo()
}
