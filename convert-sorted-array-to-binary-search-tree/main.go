package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil {
		return nil
	}
	return build(0, len(nums)-1, nums)
}

func build(start int, end int, nums []int) *TreeNode {
	if start > end {
		return nil
	}
	if start == end {
		return &TreeNode{nums[start], nil, nil}
	}
	mid := (end + start) / 2
	root := &TreeNode{nums[mid], nil, nil}
	left := build(start, mid-1, nums)
	right := build(mid+1, end, nums)
	root.Left = left
	root.Right = right
	return root
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(nums)
	fmt.Println(root.Val)
}
