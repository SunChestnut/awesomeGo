package main

import (
	"context"
	"fmt"
	"time"
)

type result struct {
	record string
	err    error
}

// search 模拟实际业务中的搜索操作
func search(term string) (string, error) {
	time.Sleep(200 * time.Millisecond)
	return term, nil
}

// process 当上下文设置的超时时间到了后，就会执行 select 语句中的第一个 case，放弃从 channel 中接收值。
// 问题在于，return 时候并没有关闭 Channel，如此可能会造成 goroutine leak
// 解决方案是，并非直接创建一个 Channel，而是创建一个 buffer channel，可以将 buffer 的值设为 1
// 直接在 return 前 close(ch) 可以解决么❓不行，关闭 channel 是停止向 channel 中发送信息的一个信号，而不能用来终止 goroutine 使用
func process(term string) {
	ch := make(chan result)
	//ch := make(chan result,1)
	go func() {
		record, err := search(term)
		ch <- result{record: record, err: err}
	}()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	go func() {
		for {
			select {
			case <-ctx.Done():
				//close(ch) ❌
				return
			case res := <-ch:
				if res.err != nil {
					return
				}
				fmt.Println("Received: ", res.record)
			}
		}
	}()
}

func main() {

}
