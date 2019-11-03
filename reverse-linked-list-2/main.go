package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	var first *ListNode
	var last *ListNode
	slice := make([]*ListNode, n-m+1)
	current := head
	for i := 0; current != nil; i++ {
		if i == m-2 && m != 1 {
			first = current
			current = current.Next
			continue
		}
		if i >= m-1 && i <= n-1 {
			slice[i-m+1] = current
		}
		current = current.Next
	}
	last = slice[len(slice)-1].Next
	if first != nil {
		first.Next = slice[len(slice)-1]
	}
	slice[0].Next = last
	for i := len(slice) - 1; i >= 1; i-- {
		slice[i].Next = slice[i-1]
	}
	if m == 1 {
		return slice[len(slice)-1]
	} else {
		return head
	}
}
