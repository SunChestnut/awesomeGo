package main

import (
	"encoding/json"
	"fmt"
)

// Order 在 Go 中字段的大小写是有含义的，但是 Go 对字段大小写的规范放在 JSON 中不适合了。因此可以在结构体的每个字段后面加上 tag，来设置该字段在 JSON 中显示的名称
// name 后面的 JSON tag 中的 omitempty 关键字表示，当构造结构体时，如果 name 为空，则不放入 name 字段到 JSON 中
type Order struct {
	ID         string `json:"id"`
	Name       string `json:"name,omitempty"`
	Quantity   int    `json:"quantity"`
	TotalPrice int    `json:"total_price"`
}

func main() {
	o := Order{
		ID: "123",
		//Name:       "Buy",
		Quantity:   888,
		TotalPrice: 9999,
	}

	// 可以直接使用 fmt.Printf("%+v\n", o) 来打印结构体
	// 关于 %v :
	// the value in a default format
	//	when printing structs, the plus flag (%+v) adds field names
	fmt.Printf("%+v\n", o)

	// ⚠️在 JSON 数据处理中，对于结构体中以小写字母开头的字段，是不会被写入 JSON 中的，因为 go 规定以小写字母开头的字段是不能被公开访问的
	marshal, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", marshal)
}
