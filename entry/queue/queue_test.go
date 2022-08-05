package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {

	q := Queue{1}

	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())

	fmt.Println(q.IsEmpty())

}

// Go 生成示例代码，该代码可在 `godoc http :6060` 开启的本地文档中查看到
func ExampleQueue_Pop() {
	q := Queue{1}

	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())

	fmt.Println(q.IsEmpty())

	// Output:
	// 1
	// 2
	// false
	// 3
	// true
}
