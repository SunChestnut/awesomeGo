package main

import (
	"awesomeGo/entry/tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postTraverse() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postTraverse()

	right := myTreeNode{myNode.node.Right}
	right.postTraverse()

	myNode.node.Print()
}

func main() {

	var root tree.Node
	root = tree.Node{Value: 3}
	// 创建了值为 0 的节点
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Traverse()

	fmt.Println()

	node := myTreeNode{&root}
	node.postTraverse()

}
