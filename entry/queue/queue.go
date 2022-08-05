package queue

// Queue An FIFO queue
type Queue []int

// Push ==>Pushes the element into the queue.
func (q *Queue) Push(v int) {
	// 为了让函数外也能看到队列中值的更改，这里的接收者应该使用指针类型
	*q = append(*q, v)
}

// Pop ==> Pops element from head.
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// IsEmpty Returns if the queue is empty of not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
