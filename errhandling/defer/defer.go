package main

import (
	"awesomeGo/functional/fib"
	"bufio"
	"errors"
	"fmt"
	"os"
)

func tryDefer() {
	// 打印结果为：3 2 1。defer 调用内部有个栈，先进后出
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 直接使用 file 写文件比较慢，可以使用 bufio.NewWriter()
	writer := bufio.NewWriter(file)
	// 需要执行 flush 操作才能将 bufio 中生成的斐波那契数值写入文件中
	defer writer.Flush()

	// 生成 20 个斐波那契数列并写入 writer 中
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func writeFileHandleError(filename string) {
	// 改成这种形式方便抛出错误，当文件已经存在时，os.O_EXCL 会报错
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	// 🍑 直接 panic(err) 的话会中断程序的运行，不推荐✖️
	//if err != nil {
	//	panic(err)
	//}

	// 🍑 打印错误信息，并返回。这种方式比较温和
	//if err != nil {
	//	fmt.Println("Error :", err)
	//	return
	//}

	// err 是一个接口，可以实现自定义类型的错误
	err = errors.New("this is a custom error")

	// 🍑 在 os.OpenFile() 函数的注释中写到：If there is an error, it will be of type *PathError.
	// 表明如果该函数报错了，那这个错误类型为 *PathError
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			// 如果 err 不是 *PathError 则直接报错
			// 什么情况下 err 不是 *PathError 类型的呢？可能在处理 err 代码的上方将 err 赋给了别的类型的错误，比如是 自定义类型的错误
			panic(err)
		} else {
			// 否则打印 *PathError 的错误信息
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
	}

	defer file.Close()

	// 直接使用 file 写文件比较慢，可以使用 bufio.NewWriter()
	writer := bufio.NewWriter(file)
	// 需要执行 flush 操作才能将 bufio 中生成的斐波那契数值写入文件中
	defer writer.Flush()

	// 生成 20 个斐波那契数列并写入 writer 中
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	//writeFile("writeFib.txt")
	writeFileHandleError("writeFib.txt")
}
