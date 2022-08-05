package mock

import "fmt"

type Retriever struct {
	Contents string
}

// Get 只要实现了 Retriever 接口中的 Get 方法，即可算是实现了 Retriever 接口
// ==> 接收者为指针类型是因为 Post 中设置了接收者为指针类型，为了统一类型而已
func (r *Retriever) Get(url string) string {
	return r.Contents
}

// Post 接收者设置为指针类型是为了更改 Contents 的值
func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

// 🍊 实现 Go 语言中常用的系统接口：String()
func (r *Retriever) String() string {
	// Sprintf formats according to a format specifier and returns the resulting string.
	return fmt.Sprintf("🌞Retriever : {Contents = %s}", r.Contents)
}
