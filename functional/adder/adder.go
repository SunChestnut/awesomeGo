package main

import "fmt"

// 【闭包】累加器
func adder() func(int) int {
	// sum 在闭包中被称为自由变量
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

/**
"正统"函数式编程
==> 不可变性：不能有状态，只有常量和函数
==> 函数只能有一个参数
*/

// 将上述累加器改为 "正统"函数
// iAdder 保存 sum 状态
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {

	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}

	a2 := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a2 = a2(i)
		fmt.Printf("0+1+2...+%d = %d\n", i, s)

	}
}
