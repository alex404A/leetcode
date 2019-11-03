package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	node  *TreeNode
	layer int
}

type Queue struct {
	list   []*Node
	length int
}

func (queue *Queue) offer(node *Node) {
	queue.list = append(queue.list, node)
	queue.length++
}

func (queue *Queue) poll() (*Node, bool) {
	if len(queue.list) == 0 {
		return nil, false
	} else {
		node := queue.list[0]
		queue.list = queue.list[1:]
		queue.length--
		return node, true
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := Queue{make([]*Node, 0), 0}
	node := &Node{root, 1}
	queue.offer(node)
	maxDepth := 1
	for queue.length > 0 {
		node, _ := queue.poll()
		maxDepth = node.layer
		if node.node.Left != nil {
			left := &Node{node.node.Left, node.layer + 1}
			queue.offer(left)
		}
		if node.node.Right != nil {
			right := &Node{node.node.Right, node.layer + 1}
			queue.offer(right)
		}
	}
	return maxDepth
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
	root := build([]int{3, 9, 20, -1, -1, 15, 7})
	result := maxDepth(root)
	fmt.Println(result)
}
