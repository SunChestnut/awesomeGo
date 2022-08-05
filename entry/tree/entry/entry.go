package main

import (
	"awesomeGo/entry/tree"
	"fmt"
)

func main() {
	node := tree.Node{Value: 3}
	node.Left = &tree.Node{Value: 0}
	node.Left.Right = &tree.Node{Value: 2}
	node.Right = &tree.Node{Value: 5}
	node.Right.Left = &tree.Node{Value: 4}

	node.Traverse()
	fmt.Println()
	node.TraverseNew()

	// 计算树中节点的个数
	nodeCount := 0
	node.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println(nodeCount)

	// 使用 Channel 遍历二叉树，并找出树中节点的最大值
	c := tree.TraverseWithChannel(&node)
	maxValue := 0
	for v := range c {
		if v.Value > maxValue {
			maxValue = v.Value
		}
	}
	fmt.Println("MaxValue = ", maxValue)
}
