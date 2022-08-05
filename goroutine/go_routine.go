package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		// 使用 go run -race xx.go 检测数据冲突
		go func(i int) {
			// 🤔 为什么要加个死循环？
			for {
				fmt.Printf("Hello form gorutine %d\n", i)
			}
		}(i)
	}
	// 休眠 1ms，防止 main 函数迅速执行完毕，而 for 循环内部还未来得及打印
	time.Sleep(time.Minute)
}
