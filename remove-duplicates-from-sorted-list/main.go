package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	first := head
	last := first
	for head != nil {
		if head.Val != last.Val {
			last.Next = head
			last = head
		}
		head = head.Next
	}
	last.Next = nil
	return first
}

func test(head *ListNode) {
	for head != nil {
		fmt.Printf("%d, ", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}

func main() {
	first := &ListNode{1, nil}
	tail := first
	tail.Next = &ListNode{1, nil}
	tail = tail.Next
	tail.Next = &ListNode{2, nil}
	tail = tail.Next
	tail.Next = &ListNode{3, nil}
	tail = tail.Next
	tail.Next = &ListNode{3, nil}
	tail = tail.Next
	result := deleteDuplicates(first)
	test(result)
}
