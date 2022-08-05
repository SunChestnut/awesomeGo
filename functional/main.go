package main

import (
	"awesomeGo/functional/fib"
	"fmt"
	"strings"
)

func main() {

	f := fib.Fibonacci()
	fmt.Println(
		f(),
		f(),
		f(),
		f(),
		f(),
	)

	fmt.Println()
	fib.PrintFileContents(f)

	s1 := "abc"
	s2 := "def"
	compare := strings.Compare(s1, s2)
	fmt.Println(compare)
}
