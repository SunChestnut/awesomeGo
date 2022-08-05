package main

import (
	"fmt"
	"io/ioutil"
)

const filename = "basic/branch/temp.txt"

func readFile() {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

// if 语句中可以赋值
func readFileBetter() {
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func calculate(a, b int, op string) int {
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
		panic("unsupported operator:" + op)
	}
	return result
}

// switch语句
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	readFile()
	readFileBetter()

	fmt.Println(
		calculate(1, 2, "+"),
		calculate(3, 5, "*"),
		calculate(0, 9, "-"),
		calculate(6, 4, "/"),
		calculate(900, 100, "a"),
	)

	fmt.Println(
		grade(0),
		grade(50),
		grade(70),
		grade(96),
		grade(100))
}
