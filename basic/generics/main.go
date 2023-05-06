package main

import "fmt"

type Slice[T int | float32 | float64] []T

type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

// MyStruct 范型类型的结构体
type MyStruct[T int | string] struct {
	Name string
	Data T
}

// IPrintData 范型接口
type IPrintData[T int | float32 | string] interface {
	Print(data T)
}

// MyChan 范型通道
type MyChan[T int | string] chan T

func main() {
	var a Slice[int] = []int{1, 2, 3}
	fmt.Printf("Type Name: %T\n", a)

	var b Slice[float32] = []float32{1.1, 2.2, 3.3}
	fmt.Printf("Type Name: %T\n", b)

	rabbit := MyMap[string, float32]{
		"number": 100,
		"weight": 3.45,
	}
	fmt.Printf("Type :%T 🍑rabbitInfo: %v\n", rabbit, rabbit)

	structA := MyStruct[string]{
		Name: "StringType",
		Data: "Car",
	}
	structB := MyStruct[int]{
		Name: "IntegerType",
		Data: 123,
	}
	fmt.Printf("Type :%T 🍑structA: %v\n", structA, structA)
	fmt.Printf("Type :%T 🍑structB: %v\n", structB, structB)

}
