package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type orderCollection struct {
	preorder []int
	inorder  []int
}

type Interval struct {
	start int
	end   int
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	collection := orderCollection{preorder, inorder}
	preInterval := Interval{0, len(preorder) - 1}
	inInterval := Interval{0, len(inorder) - 1}
	return build(preInterval, inInterval, &collection)
}

func searchRoot(interval Interval, collection *orderCollection, target int) int {
	for i := interval.start; i <= interval.end; i++ {
		if target == collection.inorder[i] {
			return i
		}
	}
	return -1
}

func build(preInterval Interval, inInterval Interval, collection *orderCollection) *TreeNode {
	if preInterval.start < 0 || preInterval.end < 0 || inInterval.start < 0 || inInterval.end < 0 {
		return nil
	}
	if preInterval.end-preInterval.start <= 1 {
		return buildSmallTree(preInterval, inInterval, collection)
	}
	root := &TreeNode{collection.preorder[preInterval.start], nil, nil}
	index := searchRoot(inInterval, collection, root.Val)
	preOffset := index - inInterval.start
	leftPreInterval := Interval{preInterval.start + 1, preInterval.start + preOffset}
	leftInInterval := Interval{inInterval.start, index - 1}
	rightPreInterval := Interval{preInterval.start + preOffset + 1, preInterval.end}
	rightInInterval := Interval{index + 1, inInterval.end}
	left := build(leftPreInterval, leftInInterval, collection)
	right := build(rightPreInterval, rightInInterval, collection)
	root.Left = left
	root.Right = right
	return root
}

func buildSmallTree(preInterval Interval, inInterval Interval, collection *orderCollection) *TreeNode {
	if preInterval.end < preInterval.start {
		return nil
	} else if preInterval.start == preInterval.end {
		return &TreeNode{collection.preorder[preInterval.start], nil, nil}
	} else if preInterval.start+1 == preInterval.end {
		root := &TreeNode{collection.preorder[preInterval.start], nil, nil}
		if collection.preorder[preInterval.start] == collection.inorder[inInterval.start] {
			root.Right = &TreeNode{collection.preorder[preInterval.end], nil, nil}
		} else {
			root.Left = &TreeNode{collection.preorder[preInterval.end], nil, nil}
		}
		return root
	} else {
		return nil
	}
}

func main() {
	preorder := []int{4, 1, 2, 3}
	inorder := []int{1, 2, 3, 4}
	root := buildTree(preorder, inorder)
	fmt.Println(root.Val)
}
