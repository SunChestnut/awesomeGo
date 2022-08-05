package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "YES我爱悠悠球!"
	fmt.Println("字节数 = ", len(s))
	fmt.Println("字符数 = ", utf8.RuneCountInString(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	// ch is a rune
	// rune 是 4 字节
	for pos, ch := range s {
		fmt.Printf("(%d %c) ", pos, ch)
	}
	fmt.Println()

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c) ", i, ch)
	}

}
