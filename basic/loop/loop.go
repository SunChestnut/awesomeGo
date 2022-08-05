package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/**
使用 for 循环，将数字转换成二进制字符
*/
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// Go 中没有 while，可使用 for 替代 while
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// fo r循环，省略全部条件时变成死循环
func forever() {
	for {
		fmt.Println("loop forever 😨😨😨")
	}
}

func main() {

	fmt.Println(
		convertToBin(5),
		convertToBin(165),
		convertToBin(888),
	)

	printFile("basic/loop/AboutNet-tools.txt")

	// `` 可以构造换行的字符串
	s := `a
b
c
d`
	printFileContents(strings.NewReader(s))

	// forever()
}
