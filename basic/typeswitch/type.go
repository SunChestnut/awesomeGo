package main

import "fmt"

/**
 * Type switches 🐱https://go.dev/tour/methods/16
 */

// 自定义类型
type intSlice []int

// interface{} 表示任何类型
func do(t interface{}) {
	// type-switch grammar
	switch v := t.(type) {
	case int:
		fmt.Printf("Current Type is %T. Twich %v is %v\n", v, v, v*2)
	case string:
		fmt.Printf("Current Type is %T. %q is %v bytes long\n", v, v, len(v))
	case intSlice:
		fmt.Printf("Current Type is %T.\n", v)
	default:
		fmt.Printf("I don't know about type %T\n", v)
	}
}

func main() {
	do(133)
	do("Hello GoLang~")
	do(true)

	slice := intSlice([]int{})
	do(slice)
}
