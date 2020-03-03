package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	list []*TreeNode
}

func (stack *Stack) length() int {
	return len(stack.list)
}

func (stack *Stack) pop() (node *TreeNode, ok bool) {
	if stack.length() == 0 {
		ok = false
		return
	}
	ok = true
	node = stack.list[len(stack.list)-1]
	stack.list = stack.list[0 : len(stack.list)-1]
	return
}

func (stack *Stack) push(node *TreeNode) {
	stack.list = append(stack.list, node)
}

type BSTIterator struct {
	stack *Stack
}

func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{&Stack{make([]*TreeNode, 0)}}
	iterator.pushAll(root)
	return iterator
}

func (this *BSTIterator) pushAll(node *TreeNode) {
	for node != nil {
		this.stack.push(node)
		node = node.Left
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	node, ok := this.stack.pop()
	if !ok {
		return -1
	}
	this.pushAll(node.Right)
	return node.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.stack.length() > 0
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
	root := build([]int{3, 1, 4, -1, 2})
	obj := Constructor(root)
	for obj.HasNext() {
		val := obj.Next()
		fmt.Println(val)
	}
}
