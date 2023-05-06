package main

import (
	"fmt"
)

// Node N 叉树
type Node struct {
	Val      int
	Children []*Node
}

// preorder N 叉树的前序遍历
func preorder(root *Node) []int {
	var res []int
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		for _, ch := range node.Children {
			fmt.Printf("current node: %v\n", ch.Val)
			dfs(ch)
		}
	}
	dfs(root)
	return res
}

func preorderTest() {
	node1 := Node{Val: 3}
	node2 := Node{Val: 2}
	node3 := Node{Val: 4}
	node4 := Node{Val: 5}
	node5 := Node{Val: 6}

	var child1 []*Node
	child1 = append(child1, &node1)
	child1 = append(child1, &node2)
	child1 = append(child1, &node3)

	root := Node{
		Val:      1,
		Children: child1,
	}

	var child2 []*Node
	child2 = append(child2, &node4)
	child2 = append(child2, &node5)

	node1.Children = child2

	res := preorder(&root)
	fmt.Println(res)
}

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrder 二叉树层序遍历
func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		var level []int
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, level)
	}
	return ans
}

func levelOrderTest() {
	root := TreeNode{Val: 3}
	node1 := &TreeNode{Val: 9}
	node2 := &TreeNode{Val: 20}
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}
	root.Left = node1
	root.Right = node2
	node2.Left = node3
	node2.Right = node4

	ans := levelOrder(&root)
	fmt.Println(ans)
}

func main() {

	preorderTest()

	levelOrderTest()
}
