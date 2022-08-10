package main

import (
	"fmt"
	"sync"
	"time"
)

/**
==> Go 内部也有针对并发原子性操作的支持，比如 `atomic.AddInt32()`。
==> 下面使用互斥量 Mutex 自己实现一个
*/

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.value++

	// 🤔 那如何给部分代码加锁解锁呢？可以使用匿名函数来实现
	//func() {
	//	a.lock.Lock()
	//	defer a.lock.Unlock()
	//
	//	a.value++
	//}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}

func main() {
	var a atomicInt
	a.increment()

	go func() {
		fmt.Println("run incrementing")
		a.increment()
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
