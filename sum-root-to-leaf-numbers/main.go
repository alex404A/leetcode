package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Result struct {
	sum   int
	layer map[int]int
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := find(root)
	return result.sum
}

func find(root *TreeNode) *Result {
	if root.Left == nil && root.Right == nil {
		return &Result{root.Val, map[int]int{1: 1}}
	}
	var left *Result
	var right *Result
	if root.Left != nil {
		left = find(root.Left)
	} else {
		left = &Result{0, make(map[int]int)}
	}
	if root.Right != nil {
		right = find(root.Right)
	} else {
		right = &Result{0, make(map[int]int)}
	}
	sum := left.sum + right.sum + add(root, left.layer) + add(root, right.layer)
	layer := merge(left.layer, right.layer)
	return &Result{sum, layer}
}

func add(node *TreeNode, layer map[int]int) int {
	result := 0
	for layer, path := range layer {
		result += path * node.Val * int(math.Pow10(layer))
	}
	return result
}

func merge(left map[int]int, right map[int]int) map[int]int {
	result := make(map[int]int)
	for layer, path := range left {
		result[layer+1] = path
	}
	for layer, path := range right {
		if _, ok := result[layer+1]; ok {
			result[layer+1] += path
		} else {
			result[layer+1] = path
		}
	}
	return result
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

func test(nums []int, expected int) {
	root := build(nums)
	actual := sumNumbers(root)
	if actual != expected {
		fmt.Printf("%v", nums)
	}
}

func main() {
	test([]int{4, 9, 0, -1, 1}, 531)
	test([]int{4, 9, 0, 5, 1}, 1026)
	test([]int{1, 2, 3}, 25)
}
