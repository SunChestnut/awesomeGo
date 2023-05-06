package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func showGoMaxProcess() {
	log.Printf("当前程序可创建的可并发执行的最大线程数量为：%v\n", runtime.GOMAXPROCS(0))
	log.Printf("cpu 个数：%v\n", runtime.NumCPU())
}

func main() {

	showGoMaxProcess()

	for i := 0; i < 1000; i++ {
		// 使用 go run -race xx.go 检测数据冲突
		go func(i int) {
			// 🤔 为什么要加个死循环？==> goroutine 启动了一个协程，该协程执行执行完毕后，便退出了。加 for 循环是为了保证该协程不结束运行，一直在从外部接收数据。
			for {
				fmt.Printf("Hello %d\n", i)
			}
		}(i)
	}
	// 休眠 1ms，防止 main 函数迅速执行完毕，而 for 循环内部还未来得及打印
	time.Sleep(time.Minute)
}
