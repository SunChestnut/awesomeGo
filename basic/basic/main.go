package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 定义方法外且作用范围在包内的变量
var aaa = 55
var bbb = "我们都是好孩子"

// 上述变量定义方式的改良版，使用var()集中定义变量
var (
	ccc = 66
	ddd = "自作聪明的孩子"
)

// 局部变量默认有初始值，整型的初始值为0，字符串的初始值为空串
func variableZeroValue() {
	var a int
	var s string
	// q -> quotation
	fmt.Printf("%d %q", a, s)
}

// 局部变量赋初始值
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "JiuJiu"
	fmt.Println(a, b, s)
}

// 省略变量类型，为局部变量赋初值
func variableTypeDeduction() {
	var a, b, c, s = 111, 222, true, "default string"
	fmt.Println(a, b, c, s)
}

// 省略变量类型并使用 : 替代var关键字，为局部变量赋初值
// PS: 用冒号替代var的方式只能在函数内使用
func variableShort() {
	a, b, c, s := 3, 4, true, "short string"
	a = 666
	fmt.Println(a, b, c, s)
}

// 最美公式——欧拉公式（ e^iπ + 1 = 0 ）
func euler() {
	// cmplx.Exp() : e的多少次方
	fmt.Printf("%0.3f", cmplx.Exp(1i*math.Pi)+1)
	fmt.Println()
}

// 强制类型转换
func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

// 常量，Go中的常量一般不会用全部大写的字符串表示，因为大写字符在Go表示某些特殊含义
const filenameInPackage = "file name in package"

func costs() {
	const filenameInFunc = "file name in func"

	fmt.Println(filenameInPackage)
	fmt.Println(filenameInFunc)
}

// Go语言中没有专指枚举类型的关键字，而是使用一组 const() 来表示
// iota 可以作为自增值的种子
func enums() {
	const (
		cpp = iota
		java
		_
		python
		ruby
	)

	fmt.Println(cpp, java, python, ruby)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello,World!")
	variableZeroValue()
	fmt.Println()
	variableInitialValue()
	variableTypeDeduction()
	variableShort()
	fmt.Println(aaa)
	fmt.Println(bbb)
	fmt.Println(ccc)
	fmt.Println(ddd)

	euler()

	triangle()

	costs()

	enums()
}
