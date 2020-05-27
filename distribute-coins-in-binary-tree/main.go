package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	move := 0
	distribute(root, &move)
	return move
}

func distribute(root *TreeNode, move *int) int {
	if root == nil {
		return 0
	}
	left := distribute(root.Left, move)
	right := distribute(root.Right, move)
	*move += abs(left) + abs(right)
	return root.Val + left + right - 1
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return 0 - a
	}
}

func main() {

}
