package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type SplitResult struct {
	sum     int
	product int64
}

func (this *SplitResult) doNothing(val int) {

}

func (this *SplitResult) calcSum(val int) {
	this.product = max(this.product, int64((this.sum-val)*val))
}

type calc func(int)

func max(a int64, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProduct(r *TreeNode) int {
	result := &SplitResult{0, 0}
	sum := dfs(r, result.doNothing)
	result.sum = sum
	dfs(r, result.calcSum)
	return int(result.product % (1e9 + 7))
}

func dfs(root *TreeNode, cb calc) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left, cb)
	right := dfs(root.Right, cb)
	sum := root.Val
	sum += left
	sum += right
	cb(sum)
	return sum
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
	root := build([]int{2, 3, 9, 10, 7, 8, 6, 5, 4, 11, 1})
	fmt.Println(maxProduct(root))
}
