package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	// 声明长度为 32，容量为 32 的切片 a
	a := make([]int, 32)
	b := a[1:16]
	fmt.Printf("len=%v, cap=%v\n", len(a), cap(a)) // output: len=32, cap=32
	fmt.Printf("len=%v, cap=%v\n", len(b), cap(b)) // output: len=15, cap=31。为何切片 b 的容量是 31 呢？

	fmt.Printf("a[2]=%v, b[1]=%v\n", a[2], b[1])

	a = append(a, 1)
	a[2] = 42
	fmt.Printf("len=%v, cap=%v\n", len(a), cap(a)) // output: len=33, cap=64。切片 a 触发了扩容，容量增大为原来的 2 倍
	fmt.Printf("a[2]=%v, b[1]=%v\n", a[2], b[1])   // output: b[2]=0
}

func Test_A(t *testing.T) {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/') // 4?
	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB

	dir1 = append(dir1, "suffix"...)     // 追加字符串到子节数组中，为何待追加的元素后有 ... 就可以了？
	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => uffixBBBB
}

func Test_C(t *testing.T) {
	str := "abcd"
	b := []byte{'a', 'b', 'c', 'd'}

	fmt.Println(reflect.DeepEqual(str, b))
}
