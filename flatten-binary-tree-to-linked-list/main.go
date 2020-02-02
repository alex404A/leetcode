package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	flat(root)
}

func flat(root *TreeNode) (head *TreeNode, tail *TreeNode) {
	if root == nil {
		head = nil
		tail = nil
		return
	}
	head = root
	tail = root
	leftHead, leftTail := flat(root.Left)
	rightHead, rightTail := flat(root.Right)
	root.Left = nil
	root.Right = nil
	if leftHead != nil {
		head.Right = leftHead
		tail = leftTail
	}
	if rightHead != nil {
		tail.Right = rightHead
		tail = rightTail
	}
	return
}

func build(array []int) *TreeNode {
	if len(array) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, 1)
	nodes[0] = &TreeNode{array[0], nil, nil}
	for i := 1; i < len(array); i++ {
		val := array[i]
		if val == -1 {
			continue
		}
		node := &TreeNode{val, nil, nil}
		nodes = append(nodes, node)
		j := i / 2
		k := i % 2
		if k == 1 {
			parent := nodes[j]
			parent.Left = node
		}
		if k == 0 {
			parent := nodes[j-1]
			parent.Right = node
		}
	}
	return nodes[0]
}

func main() {
	nums := []int{1, 2, 5, 3, 4, -1, 6}
	root := build(nums)
	flatten(root)
	fmt.Println(nums)
}
