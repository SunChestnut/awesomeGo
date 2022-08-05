package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("s = %v,len = %d, cap = %d\n", s, len(s), cap(s))
}

func operateSlices() {
	// 创建slice方式一
	var s1 []int

	for i := 0; i < 100; i++ {
		printSlice(s1)
		s1 = append(s1, 2*i+1)
	}

	fmt.Println(s1)

	// 创建slice方式二：声明并直接赋值
	var s2 = []int{2, 4, 6, 8, 10}
	printSlice(s2)

	// 创建slice方式三：设置slice大小
	var s3 = make([]int, 16)
	printSlice(s3)

	// 创建slice方式四：设置slice大小并设置cap
	var s4 = make([]int, 10, 32)
	printSlice(s4)

	// 拷贝slice
	copy(s3, s2)
	printSlice(s3)

	// 删除slice中的某个元素，比如删除s3中的元素8
	fmt.Println("Deleting element from slice")
	s3 = append(s3[:3], s3[4:]...)
	printSlice(s3)

	// 弹出slice中的头元素
	fmt.Println("Popping from front...")
	s3 = s3[1:]
	printSlice(s3)

	// 弹出slice中的尾元素
	fmt.Println("Popping from back...")
	s3 = s3[:len(s3)-1]
	printSlice(s3)

	fmt.Println("--------")

	var slichaha = make([]int, 0)
	fmt.Println(slichaha)

	for i := 0; i <= 20; i = i + 2 {
		slichaha = append(slichaha, i)
	}

	fmt.Println(slichaha)

	var value = slichaha[0]
	slichaha = slichaha[1:]

	fmt.Println(slichaha)
	fmt.Println(value)
}
