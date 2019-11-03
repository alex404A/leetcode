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
	if head.Next == nil {
		return head
	}
	cur := &ListNode{-2, nil}
	cnt := 0
	first := &ListNode{-1, nil}
	last := first
	for head != nil {
		if head.Val != cur.Val {
			if cnt == 1 {
				last.Next = cur
				last = cur
			}
			cur = head
			cnt = 1
		} else {
			cnt++
		}
		head = head.Next
	}
	if cnt == 1 {
		last.Next = cur
		last = cur
	}
	last.Next = nil
	return first.Next
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
	tail.Next = &ListNode{3, nil}
	tail = tail.Next
	tail.Next = &ListNode{3, nil}
	// tail = tail.Next
	// tail.Next = &ListNode{4, nil}
	// tail = tail.Next
	// tail.Next = &ListNode{4, nil}
	// tail = tail.Next
	// tail.Next = &ListNode{5, nil}
	// tail = tail.Next
	result := deleteDuplicates(first)
	test(result)
}
