package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
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
}
