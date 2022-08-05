package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func div(a, b int) (q, r int) {
	//return a / b, a % b
	q = a / b
	r = a % b
	return
}

func cal(p, q int) (int, int, int) {
	return p - q, p * q, p + q
}

func eval(a, b int, op string) (int, error) {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		return 0, fmt.Errorf("unsupported operator: %s", op)
	}
	return result, nil
}

/**
【函数式编程】
	op() 为一个包括两个 int 型参数的函数，a、b 为两个 int 型的传入 apply() 函数的参数，apply() 函数返回类型为 int
*/
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d,%d)\n", opName, a, b)
	return op(a, b)
}

func sum(numbers ...int) int {
	res := 0
	for i := range numbers {
		res += numbers[i]
	}
	return res
}

func main() {

	q, r := div(14, 3)
	fmt.Println(q, r)

	fmt.Println("---------------------------------------")

	if result, err := eval(3, 4, "a"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	// 匿名函数作为参数
	fmt.Println(
		apply(func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println("---------------------------------------")

}
