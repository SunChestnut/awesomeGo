package main

import (
	"awesomeGo/retriever/mock"
	"awesomeGo/retriever/real"
	"fmt"
	"time"
)

// Retriever 接口，内部包含一 Get 接口，可以接收文件
type Retriever interface {
	Get(url string) string
}

// Poster 接口，发送
type Poster interface {
	Post(url string, form map[string]string) string
}

// RetrieverPoster 将 Retriever 和 Poster 接口组合起来使用
type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "https://studygolang.com/"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{"name": "abc", "course": "golang"})
}

// 使用组合接口，就可以同时调用两个接口中的方法
// ==> 如果想同时调用的接口不多的话，也可以直接写成这样
//func connect(r Retriever, p Poster) {
//	r.Get("")
//	p.Post("", map[string]string{})
//}
func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{"contents": "🍬Another faked studygolang.com"})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("Inspecting...", r)
	// T for type, v for value
	fmt.Printf("==> %T %v\n", r, r)
	fmt.Printf("==> Type Switch: ")

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}

	fmt.Println()
}

func otherInspect(r Retriever) {
	// Type assertion: 通过 .() 中存放类型的名字来获取真正的类型
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.UserAgent)

	// 严格版本可以加上对结果是否 ok 的判断
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("💔Not a mock retriever")
	}
}

func main() {
	var r Retriever
	retriever := mock.Retriever{Contents: "🤔This is a fake studygolang.com"}
	r = &retriever
	inspect(r)

	fmt.Println("Try a session")
	fmt.Println(session(&retriever) + "\n")

	// 如果直接写成 r = real.Retriever{} 则会报错：real.Retriever does not implement Retriever (Get method has pointer receiver)
	// 表明 real.Retriever 实现 Retriever 中的 Get() 方法时，接收者是指针类型
	// 所以这里构造接收者时应该取地址
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

}
