package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	max := 0 - int(^uint(0)>>1) - 1
	check(root, &max)
	return max
}

func check(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := getMax(0, check(root.Left, max))
	right := getMax(0, check(root.Right, max))
	*max = getMax(*max, left+right+root.Val)
	return getMax(left, right) + root.Val
}

func getMax(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	root := &TreeNode{-7, nil, nil}
	left := &TreeNode{-5, nil, nil}
	root.Left = left
	result := maxPathSum(root)
	fmt.Println(result)
}
