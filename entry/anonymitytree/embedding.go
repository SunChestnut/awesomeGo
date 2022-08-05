package main

import (
	"awesomeGo/entry/tree"
	"fmt"
)

type withoutNameTree struct {
	*tree.Node
}

func (node *withoutNameTree) preTraverse() {
	if node == nil || node.Node == nil {
		return
	}
	fmt.Print(node.Value, " ")
	left := withoutNameTree{node.Left}
	left.preTraverse()
	right := withoutNameTree{node.Right}
	right.preTraverse()

}

func main() {

	root := withoutNameTree{&tree.Node{Value: 1}}
	root.Left = &tree.Node{Value: 2}
	root.Left.Left = &tree.Node{Value: 4}
	root.Left.Right = &tree.Node{Value: 5}
	root.Right = &tree.Node{Value: 3}
	root.Right.Left = &tree.Node{Value: 6}
	root.Right.Right = &tree.Node{Value: 7}

	root.preTraverse()

	fmt.Println()
	root.Traverse()
}
