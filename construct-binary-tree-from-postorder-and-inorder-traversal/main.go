package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type orderCollection struct {
	postorder []int
	inorder   []int
}

type Interval struct {
	start int
	end   int
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	collection := orderCollection{postorder, inorder}
	postInterval := Interval{0, len(postorder) - 1}
	inInterval := Interval{0, len(inorder) - 1}
	return build(postInterval, inInterval, &collection)
}

func searchRoot(interval Interval, collection *orderCollection, target int) int {
	for i := interval.start; i <= interval.end; i++ {
		if target == collection.inorder[i] {
			return i
		}
	}
	return -1
}

func build(postInterval Interval, inInterval Interval, collection *orderCollection) *TreeNode {
	if postInterval.start < 0 || postInterval.end < 0 || inInterval.start < 0 || inInterval.end < 0 {
		return nil
	}
	if postInterval.end-postInterval.start <= 1 {
		return buildSmallTree(postInterval, inInterval, collection)
	}
	root := &TreeNode{collection.postorder[postInterval.end], nil, nil}
	index := searchRoot(inInterval, collection, root.Val)
	postOffset := index - inInterval.start
	leftPostInterval := Interval{postInterval.start, postInterval.start + postOffset - 1}
	leftInInterval := Interval{inInterval.start, index - 1}
	rightPostInterval := Interval{postInterval.start + postOffset, postInterval.end - 1}
	rightInInterval := Interval{index + 1, inInterval.end}
	left := build(leftPostInterval, leftInInterval, collection)
	right := build(rightPostInterval, rightInInterval, collection)
	root.Left = left
	root.Right = right
	return root
}

func buildSmallTree(postInterval Interval, inInterval Interval, collection *orderCollection) *TreeNode {
	if postInterval.end < postInterval.start {
		return nil
	} else if postInterval.start == postInterval.end {
		return &TreeNode{collection.postorder[postInterval.start], nil, nil}
	} else if postInterval.start+1 == postInterval.end {
		root := &TreeNode{collection.postorder[postInterval.end], nil, nil}
		if collection.postorder[postInterval.start] == collection.inorder[inInterval.start] {
			root.Left = &TreeNode{collection.postorder[postInterval.start], nil, nil}
		} else {
			root.Right = &TreeNode{collection.postorder[postInterval.start], nil, nil}
		}
		return root
	} else {
		return nil
	}
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree(inorder, postorder)
	fmt.Println(root.Val)
}
