package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	list []*TreeNode
}

func (this *Stack) push(node *TreeNode) {
	this.list = append(this.list, node)
}

func (this *Stack) pop() (node *TreeNode, ok bool) {
	if len(this.list) == 0 {
		ok = false
		return
	} else {
		node = this.list[len(this.list)-1]
		this.list = this.list[:len(this.list)-1]
		ok = true
		return
	}
}

func kthSmallest(root *TreeNode, k int) int {
	leftStack := &Stack{make([]*TreeNode, 0)}
	node := root
	for node != nil {
		leftStack.push(node)
		node = node.Left
	}
	for len(leftStack.list) > 0 {
		node, _ := leftStack.pop()
		total := calTotalFromRight(node)
		if total >= k {
			return findFromRight(node, k)
		} else {
			k -= total
		}
	}
	if k == 1 {
		return root.Val
	} else {
		return kthSmallest(root.Right, k-1)
	}
}

func calTotalFromRight(node *TreeNode) int {
	total := 1
	start := node.Right
	if start != nil {
		total += calTotal(start)
	}
	return total
}

func calTotal(node *TreeNode) int {
	total := 1
	if node.Left != nil {
		total += calTotal(node.Left)
	}
	if node.Right != nil {
		total += calTotal(node.Right)
	}
	return total
}

func findFromRight(node *TreeNode, k int) int {
	if k == 1 {
		return node.Val
	}
	all := []*TreeNode{node}
	nodes := make([]*TreeNode, 0)
	if node.Right != nil {
		nodes = append(nodes, node.Right)
	}
	for len(nodes) > 0 {
		node := nodes[0]
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
		all = append(all, node)
		nodes = nodes[1:]
	}
	sort.SliceStable(all, func(i, j int) bool {
		return all[i].Val <= all[j].Val
	})
	return all[k-1].Val
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
	// root := build([]int{5, 3, 6, 2, 4, -1, -1, 1})
	// result := kthSmallest(root, 3)
	// root := build([]int{3, 1, 4, -1, 2})
	// result := kthSmallest(root, 1)
	root := build([]int{5, 3, 6, 2, 4, -1, -1, 1})
	result := kthSmallest(root, 6)
	fmt.Println(result)
}
