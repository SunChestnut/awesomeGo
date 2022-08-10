package main

import (
	"encoding/json"
	"fmt"
)

/**
==> 使用 struct 构建复杂的 JSON
*/

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	ID         string    `json:"id"`
	Item       OrderItem `json:"item"`
	Quantity   int       `json:"quantity"`
	TotalPrice int       `json:"total_price"`
}

type Cart struct {
	ID         string      `json:"id"`
	Items      []OrderItem `json:"items"`
	TotalPrice int         `json:"total_price"`
}

func marshalOrder() {
	o := Order{
		ID: "23333",
		Item: OrderItem{
			ID:    "0001",
			Name:  "lesson-1",
			Price: 788,
		},
		Quantity:   2,
		TotalPrice: 1576,
	}

	value, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", value)
}

func marshalCart() []byte {
	cart := Cart{
		ID: "9988",
		Items: []OrderItem{
			{
				ID:    "1",
				Name:  "skirt",
				Price: 199,
			},
			{
				ID:    "2",
				Name:  "shoes",
				Price: 344,
			},
		},
		TotalPrice: 500,
	}

	value, err := json.Marshal(cart)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", value)
	return value
}

// 从 JSON 中解析数据
func unMarshal(content []byte, cart *Cart) {
	err := json.Unmarshal(content, cart)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", cart)
}

func main() {
	marshalOrder()

	content := marshalCart()
	var c Cart
	unMarshal(content, &c)
}
