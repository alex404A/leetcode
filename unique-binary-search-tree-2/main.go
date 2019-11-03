package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return make([]*TreeNode, 0)
	}
	return generate(1, n)
}

func generate(start int, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	if start == end {
		return []*TreeNode{&TreeNode{start, nil, nil}}
	}

	list := make([]*TreeNode, 0)

	for i := start; i <= end; i++ {
		left := generate(start, i-1)
		right := generate(i+1, end)
		for _, lnode := range left {
			for _, rnode := range right {
				root := TreeNode{i, lnode, rnode}
				list = append(list, &root)
			}
		}
	}

	return list
}

func print(root *TreeNode) {
	queue := []*TreeNode{root}
	list := make([]int, 0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			list = append(list, -1)
		} else {
			list = append(list, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	fmt.Println(list)
}

func main() {
	list := generateTrees(3)
	for _, root := range list {
		print(root)
	}
}
