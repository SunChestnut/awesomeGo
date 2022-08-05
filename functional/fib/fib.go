package fib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Fibonacci 斐波那契数列生成器
func Fibonacci() IntGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// IntGen 类型均可以实现接口，所以想为 fibonacci() 函数实现接口，需要先定义类型
type IntGen func() int

// 🤯Go 语言特色功能：函数也能实现接口，接收者和函数的普通参数类似，只不过调用函数时，接收者和函数参数所放的位置不同
func (g IntGen) Read(p []byte) (n int, err error) {
	//  取得斐波那契数列的下一个值
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	// 将值写入到字节数组中
	// TODO : incorrect if p is too small !
	return strings.NewReader(s).Read(p)
}

func PrintFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
