package main

import (
	"errors"
	"fmt"
)

// recover : 仅在 defer 函数调用中使用，可以获取 panic 的值，如果无法处理，可以重新 panic
func tryRecover() {
	// 匿名函数后加 () 表示：声明完函数后直接调用该函数
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

func tryRecoverAgain(n int) {
	// PS: recover 一定要写在 panic 的上面，否则还没执行到 recover 呢，程序就先退出了
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred: ", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do : %v", err))
		}
	}()

	switch {
	case n < 0:
		panic("You Suck")
	case n < 60:
		panic(123)
	case n > 90:
		panic(errors.New("you rock"))
	}
}

func check(n int) {
	// 只供学习 recover 使用，实际开发中不推荐这种将 panic 和 recover 写成 try...catch 的形式
	defer func() {
		if recover() != nil {
			fmt.Println("is neither")
		}
	}()

	if n < 0 {
		fmt.Println("is negative")
	} else if n > 0 {
		fmt.Println("is positive")
	} else {
		panic("undefined")
	}
}

func main() {
	//tryRecover()

	//tryRecoverAgain(-19) // 当 n < 0 时，panic 函数为字符串，不属于 error 类型，因此在 recover 中会被重新 panic 出去
	//tryRecoverAgain(100) // 当 n>90 时，代码中创建了一个 error，且在 defer 中捕获了 error，因此执行当前代码不会发生 panic

	check(2)   // is positive
	check(-90) // is negative
	check(0)   // no panic,just out put 'is neither' on terminal

	//var a interface{}
	//a = errors.New("some error is coming")
	//err := a.(error)
	//fmt.Println(err)
}
