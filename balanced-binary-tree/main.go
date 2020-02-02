package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	result, _ := check(root)
	return result
}

func check(root *TreeNode) (bool, int) {
	if root.Left != nil && root.Right != nil {
		leftResult, leftDepth := check(root.Left)
		rightResult, rightDepth := check(root.Right)
		return isLessThanOne(leftDepth, rightDepth) && leftResult && rightResult, max(leftDepth, rightDepth) + 1
	} else if root.Left == nil && root.Right == nil {
		return true, 1
	} else {
		if root.Left == nil {
			return root.Right.Left == nil && root.Right.Right == nil, 2
		} else {
			return root.Left.Left == nil && root.Left.Right == nil, 2
		}
	}

}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func isLessThanOne(a int, b int) bool {
	if a >= b {
		return a-b <= 1
	} else {
		return b-a <= 1
	}
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

func test(nums []int, expected bool) {
	root := build(nums)
	actual := isBalanced(root)
	if actual != expected {
		fmt.Printf("%v", nums)
	}
}

func main() {
	test([]int{3, 9, 20, -1, -1, 15, 7}, true)
	test([]int{1, 2, 2, 3, 3, -1, -1, 4, 4}, false)
	test([]int{1, 2}, true)
	test([]int{1, 1, 2, -1, 2, 2}, true)
}
