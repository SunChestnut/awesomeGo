package main

import (
	"fmt"
)

// recover : 仅在 defer 函数调用中使用，可以获取 panic 的值，如果无法处理，可以重新 panic
func tryRecover() {
	// ❓ 匿名函数后加 () 表示：声明完函数后直接调用该函数
	defer func() {

		r := recover()
		fmt.Printf("%T\n", r)

		if err, ok := r.(error); ok {
			fmt.Println("Error occurred: ", err)
		} else {
			// 当 recover 接收的 panic 不是 error 类型时，就会重新 panic
			panic(fmt.Sprintf("I don't know what to do : %v", err))
		}
	}()

	//panic(errors.New("this is an error "))
	panic(123)
}

func main() {
	tryRecover()

	//var a interface{}
	//a = errors.New("some error is coming")
	//err := a.(error)
	//fmt.Println(err)
}
