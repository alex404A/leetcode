package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	result := false
	result = hasPathSum(root.Left, sum-root.Val)
	if result {
		return true
	}
	result = hasPathSum(root.Right, sum-root.Val)
	if result {
		return true
	}
	return false
}

func main() {

}
