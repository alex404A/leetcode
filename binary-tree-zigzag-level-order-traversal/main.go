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

type Stack struct {
	list     [][]*Node
	curLayer int
	length   int
}

func (stack *Stack) offer(node *Node) {
	if len(stack.list) < node.layer {
		stack.list = append(stack.list, make([]*Node, 0))
	}
	stack.list[node.layer-1] = append(stack.list[node.layer-1], node)
	stack.length++
}

func (stack *Stack) poll() (*Node, bool) {
	if stack.length == 0 {
		return nil, false
	} else {
		curStack := stack.list[stack.curLayer-1]
		node := curStack[len(curStack)-1]
		stack.list[stack.curLayer-1] = curStack[:len(curStack)-1]
		stack.length--
		if len(stack.list[stack.curLayer-1]) == 0 {
			stack.curLayer++
		}
		return node, true
	}
}

type Container struct {
	list [][]int
}

func (container *Container) add(node *Node) {
	if node.layer > len(container.list) {
		container.list = append(container.list, make([]int, 0))
	}
	container.list[node.layer-1] = append(container.list[node.layer-1], node.node.Val)
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}
	container := Container{make([][]int, 1)}
	stack := Stack{make([][]*Node, 0), 1, 0}
	node := &Node{root, 1}
	stack.offer(node)
	for stack.length > 0 {
		node, _ := stack.poll()
		container.add(node)
		var left *Node
		var right *Node
		if node.node.Left != nil {
			left = &Node{node.node.Left, node.layer + 1}
		}
		if node.node.Right != nil {
			right = &Node{node.node.Right, node.layer + 1}
		}
		if node.layer%2 == 1 {
			if left != nil {
				stack.offer(left)
			}
			if right != nil {
				stack.offer(right)
			}
		} else {
			if right != nil {
				stack.offer(right)
			}
			if left != nil {
				stack.offer(left)
			}
		}
	}
	return container.list
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
	root := build([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	result := zigzagLevelOrder(root)
	fmt.Println(result)
}
