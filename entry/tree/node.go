package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	node.Value = value
}

func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

// TraverseNew 将上述只能中序遍历打印二叉树的 Traverse() 函数进行改造，让其可以实现更多的功能
func (node *Node) TraverseNew() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(node *Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func TraverseWithChannel(node *Node) chan *Node {
	c := make(chan *Node)
	go func(node *Node) {
		node.TraverseFunc(func(node *Node) {
			fmt.Println("Value = ", node.Value)
			c <- node
		})
		close(c)
	}(node)
	return c
}
