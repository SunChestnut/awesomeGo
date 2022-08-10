package main

import (
	"encoding/json"
	"fmt"
)

func provideJson() string {
	return `{
    "data": [
        {
            "synonym": "",
            "weight": "0.6",
            "word": "真丝",
            "tag": "材质"
        },
        {
            "synonym": "",
            "weight": "0.8",
            "word": "韩都衣舍",
            "tag": "品牌"
        },
        {
            "synonym": "连身裙;联衣裙",
            "weight": "1.0",
            "word": "连衣裙",
            "tag": "品类"
        }
    ]
}`
}

func unmarshalWithMap(res string) {
	var m map[string]any
	err := json.Unmarshal([]byte(res), &m)
	if err != nil {
		panic(err)
	}

	fmt.Println(m["data"])
	fmt.Println(m["data"].([]any)[2])
	fmt.Println(m["data"].([]any)[2].(map[string]any)["synonym"])
}

func unmarshalWithStruct(res string) {
	m := struct {
		Data []struct {
			Synonym string `json:"synonym"`
			Weight  string `json:"weight"`
			Word    string `json:"word"`
			Tag     string `json:"tag"`
		}
	}{}
	err := json.Unmarshal([]byte(res), &m)
	if err != nil {
		panic(err)
	}

	fmt.Println(m.Data[2].Synonym)
}

func main() {
	unmarshalWithMap(provideJson())
	unmarshalWithStruct(provideJson())
}
