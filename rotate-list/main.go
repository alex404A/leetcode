package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	listLen := getLen(head)
	reverse := buildSlice(head, listLen)
	remain := k % listLen
	fmt.Println(listLen)
	for i := 0; i < remain; i++ {
		tail := reverse[listLen-1-i]
		tailSec := reverse[listLen-2-i]
		tail.Next = head
		tailSec.Next = nil
		head = tail
	}
	return head
}

func getLen(head *ListNode) int {
	len := 1
	for head.Next != nil {
		len++
		head = head.Next
	}
	return len
}

func buildSlice(head *ListNode, listLen int) []*ListNode {
	result := make([]*ListNode, listLen)
	for i := 0; i < listLen; i++ {
		result[i] = head
		head = head.Next
	}
	return result
}

func main() {
	node1 := ListNode{1, nil}
	node2 := ListNode{2, nil}
	node3 := ListNode{3, nil}
	node1.Next = &node2
	node2.Next = &node3
	rotateRight(&node1, 4)
}
