package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	firstHead := &ListNode{-1, nil}
	secondHead := &ListNode{-1, nil}
	firstTail := firstHead
	secondTail := secondHead
	for head != nil {
		if head.Val < x {
			firstTail.Next = head
			firstTail = head
		} else {
			secondTail.Next = head
			secondTail = head
		}
		head = head.Next
	}
	firstTail.Next = nil
	secondTail.Next = nil
	if firstHead.Next == nil {
		return secondHead.Next
	} else {
		result := firstHead.Next
		firstTail.Next = secondHead.Next
		return result
	}
}

func main() {

}
