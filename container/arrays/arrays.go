package main

import "fmt"

// Go 中的数组是值类型
func printArray(arr *[6]string) {
	arr[0] = "哈士奇"
	for i := range arr {
		fmt.Println(arr[i])
	}
}

func main() {

	// 数组：数量写在类型前
	var arr1 [5]int

	arr2 := [3]int{1, 3, 5}

	// 不声明数组元素个数，让编译器来数一数，二四六七八～
	arr3 := [...]int{2, 4, 6, 8, 10}

	// 二维数组
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3, grid)

	// 原始数组的遍历方式
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	// 遍历数组方式一：通过下标获取元素值
	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	// 遍历数组方式二：直接获取元素值
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	// 遍历数组方式三：不要下标
	for _, v := range arr3 {
		fmt.Println(v)
	}

	dogs := [...]string{"吉娃娃", "博美", "马尔济斯犬", "约克夏梗", "贵宾犬", "蝴蝶犬"}
	// 这里要改写成 取地址 的方式
	printArray(&dogs)
	fmt.Printf("Let's the first dog out : %s", dogs[0])
}
