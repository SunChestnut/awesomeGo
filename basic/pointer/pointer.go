package main

// Go只有值传递，如果参数为指针，则会将传入的变量的引用拷贝一份作为参数传入
// 交换两个变量的值可以使用指针
func swap(a, b *int) {
	*a, *b = *b, *a
}

// 交换两个变量的值除了使用指针，还可以将交换的结果直接返回回去
func swapAnother(a, b int) (i, j int) {
	return b, a
}

func main() {

}
