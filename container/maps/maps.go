package main

import (
	"fmt"
	"sort"
)

// Map 是 Hashmap，无序，如果想排序，可以将 key 放到 slice 中，slice 是有序的
func main() {

	// 创建 map 方式一
	m := map[string]string{
		"name":    "虎弟和啾啾",
		"course":  "捕鼠教程",
		"site":    "www.不爱看就滚蛋.com",
		"quality": "paw",
	}

	// 创建 map 方式二，m2 == empty map
	m2 := make(map[string]int)

	// 创建 map 方式三，m3 == nil
	var m3 map[string]int

	fmt.Println(m, m2, m3)

	// 遍历 map
	fmt.Println("Traveling map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 根据 key 获取 value
	fmt.Println("Getting values")
	valueA, ok := m["name"]
	fmt.Println(valueA, ok)

	if valueB, ok := m["course"]; ok {
		fmt.Println(valueB)
	} else {
		fmt.Println("key does not exist.")
	}

	// 删除 map 中的元素
	fmt.Println("deleting map")
	quality, ok := m["quality"]
	fmt.Println(quality, ok)
	delete(m, "quality")
	quality, ok = m["quality"]
	fmt.Println(quality, ok)

	sliceSort := []int{7, 9, 4, 55, 0, -1, 3}
	sort.Ints(sliceSort)
	fmt.Println(sliceSort)
}
