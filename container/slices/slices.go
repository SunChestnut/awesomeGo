package main

import "fmt"

func updateSlice(arr []int) {
	arr[0] = 999
}

func sliceHigher() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]

	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
}

func addToSlice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := arr[3:5]
	s3 := append(s2, 10)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

func sliceTest() {
	fmt.Println("----------------")

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}

	sFirst := arr[:6]
	fmt.Println("sFirst = ", sFirst)
	sFirst[1] = 999
	fmt.Println("sFirst change = ", sFirst)
	fmt.Println("after sFirst change, array = ", arr)

	sSecond := sFirst[0:2]
	fmt.Println("sSecond = ", sSecond)
	sSecond[1] = 6666
	fmt.Println("sSecond change = ", sSecond)
	fmt.Println("after sSecond change, array = ", arr)

	fmt.Println("----------------")
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])

	s1 := arr[2:]
	fmt.Println("s1 = ", s1)

	s2 := arr[:]
	fmt.Println("arr[:]", s2)

	fmt.Print("After update slice s1 : ")
	updateSlice(s1)
	fmt.Println(s1)

	fmt.Print("After update slice s2 : ")
	updateSlice(s2)
	fmt.Println(s2)

	sliceHigher()

	sliceTest()

	addToSlice()
}
