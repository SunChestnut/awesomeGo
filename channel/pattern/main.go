package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
==> 【并发编程模式 ——> 变现】消息生成器
*/

func msgGen(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			// 为防止写入消息过快不方便阅读，设置 [0...2]s 的随机休眠时间
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("service-%s message %d", name, i)
			i++
		}
	}()
	return c
}

// 将 c1 和 c2 管道中的值输入到 c 中，再使用 c 统一将值输出出去
func funIn(c1, c2 chan string) chan string {
	c := make(chan string)

	go func() {
		for {
			// 将 c1 中的数据取出来，写入 c
			c <- <-c1
		}
	}()

	go func() {
		for {
			c <- <-c2
		}
	}()

	return c
}

func funInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case n := <-c1:
				c <- n
			case n := <-c2:
				c <- n
			}
		}
	}()
	return c
}

// 当参数为多个 channel 时，如何将多个 channel 中的值全部接收到一个 channel 中
func funInWithMultiChan(chs ...chan string) chan string {
	c := make(chan string)
	for _, ch := range chs {
		go func(in chan string) {
			for {
				c <- <-in
			}
		}(ch)
	}
	return c
}

// 【非阻塞等待】使用 select ➕default 创建非阻塞的 channel
func nonBlockingWait(c chan string) (string, bool) {
	select {
	case n := <-c:
		return n, true
	default:
		return "", false
	}
}

//【超时机制】
func timeoutWait(c chan string, timeout time.Duration) (string, bool) {
	select {
	case n := <-c:
		return n, true
	case <-time.After(timeout):
		return "", false
	}
}

func msgGenWithExist(name string, done chan struct{}) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(5000)) * time.Millisecond):
				c <- fmt.Sprintf("service %s: message %d", name, i)
			case <-done:
				fmt.Println("cleaning up")
				return
			}
			i++
		}
	}()
	return c
}

func consumeMsg() {
	done := make(chan struct{})
	c := msgGenWithExist("abc", done)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	done <- struct{}{}
}

func main() {

	//c1 := msgGen("a")
	//c2 := msgGen("b")

	//c := funIn(c1, c2)
	//c := funInBySelect(c1, c2)
	//c := funInWithMultiChan(c1, c2)
	//for {
	//	fmt.Println(<-c)
	//}

	//fmt.Println("Non-Blocking test")
	//for {
	//	fmt.Println(<-c1)
	//	if value, ok := nonBlockingWait(c2); ok {
	//		fmt.Println(value)
	//	} else {
	//		fmt.Println("no message from service-b")
	//	}
	//}

	//fmt.Println("Timeout test")
	//for {
	//	fmt.Println(<-c2)
	//	if value, ok := timeoutWait(c1, time.Duration(2*time.Second)); ok {
	//		fmt.Println(value)
	//	} else {
	//		fmt.Println("timeout")
	//	}
	//}

	consumeMsg()

	time.Sleep(time.Second)
}
