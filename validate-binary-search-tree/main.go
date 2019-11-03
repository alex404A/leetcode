package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func isValidBST(root *TreeNode) bool {
	status, _, _ := isValid(root)
	return status
}

func isValid(root *TreeNode) (status bool, leftMmin int, rightMax int) {
	if root == nil {
		return true, MaxInt, MinInt
	}
	if root.Left == nil && root.Right == nil {
		return true, root.Val, root.Val
	}
	if root.Left != nil && root.Val <= root.Left.Val {
		return false, MaxInt, MinInt
	}
	if root.Right != nil && root.Val >= root.Right.Val {
		return false, MaxInt, MinInt
	}
	status, leftMmin1, rightMax1 := isValid(root.Left)
	if status == false || rightMax1 >= root.Val {
		return false, MaxInt, MinInt
	}
	status, leftMmin2, rightMax2 := isValid(root.Right)
	if status == false || leftMmin2 <= root.Val {
		return false, MaxInt, MinInt
	}
	return true, getMin(leftMmin1, root.Val), getMax(rightMax2, root.Val)
}

func getMin(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func getMax(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	root := TreeNode{3, nil, nil}
	left := TreeNode{1, nil, nil}
	right := TreeNode{5, nil, nil}
	root.Left = &left
	root.Right = &right
	left1 := TreeNode{0, nil, nil}
	right1 := TreeNode{2, nil, nil}
	root.Left.Left = &left1
	root.Left.Right = &right1
	left2 := TreeNode{4, nil, nil}
	right2 := TreeNode{6, nil, nil}
	root.Right.Left = &left2
	root.Right.Right = &right2
	root.Left.Right.Right = &TreeNode{3, nil, nil}
	isValidBST(&root)
}
