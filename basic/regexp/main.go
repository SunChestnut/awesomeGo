package main

import (
	"fmt"
	"regexp"
)

const simpleText = "My email is yuyu21019@gmail.com@123abc"

const complexText = `
My email is yuyu21019@gmail.com@123abc
email1 is abc@163.com
email2 is def@qq.com
email3 is abc@def.cn.com
`

// findSingleEmail 匹配一个简单文本中的单个 Email
func findSingleEmail(text string) {
	// 当你非常确信你的正则不会出错时，可以使用 regexp.MustCompile()
	//re, err := regexp.Compile("yuyu21019@gmail.com")

	// . 表示匹配任何一个字符，+ 表示前面的字符出现了一次或者多次
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9.]+`)
	match := re.FindString(text)
	fmt.Println(match)
}

// findMultiEmail 匹配一个复杂文本中所有的 Email
func findMultiEmail(text string) {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9.]+`)
	// FindAllString() 表示查找字符串中所有匹配正则的子串，第二个参数 n 表示要查找多少个，-1 表示要查找所有
	match := re.FindAllString(text, -1)
	fmt.Println(match)
}

// 匹配一个复杂文本中所有邮箱相关的关键字，包括：邮箱、登录名、域名和子域名
func findInfos(text string) {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, v := range match {
		fmt.Println(v)
	}
}

func main() {
	findSingleEmail(simpleText)
	findMultiEmail(complexText)
	findInfos(complexText)
}
