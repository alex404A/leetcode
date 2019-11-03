package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type result struct {
	values []int
}

func (result *result) push(v int) {
	result.values = append(result.values, v)
}

type stack struct {
	values []*TreeNode
}

func (s *stack) push(node *TreeNode) {
	s.values = append(s.values, node)
}

func (s *stack) pop() (*TreeNode, bool) {
	if len(s.values) > 0 {
		result := s.values[len(s.values)-1]
		s.values = s.values[:len(s.values)-1]
		return result, true
	} else {
		return nil, false
	}
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	s := stack{make([]*TreeNode, 0, 16)}
	result := result{make([]int, 0)}
	findLeftMost(root, &s)
	var node *TreeNode
	for len(s.values) > 0 {
		node, _ = s.pop()
		result.push(node.Val)
		if node.Right != nil {
			findLeftMost(node.Right, &s)
		}
	}
	return result.values
}

func findLeftMost(node *TreeNode, s *stack) {
	for node.Left != nil {
		s.push(node)
		node = node.Left
	}
	s.push(node)
}

func main() {
	root := TreeNode{1, nil, nil}
	// node1 := TreeNode{2, nil, nil}
	// node2 := TreeNode{3, nil, nil}
	// root.Right = &node1
	// node1.Left = &node2
	result := inorderTraversal(&root)
	fmt.Println(result)
}
