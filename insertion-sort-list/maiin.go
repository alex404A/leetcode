package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	sortedList := make([]*ListNode, 1)
	sortedList[0] = head
	cur := head.Next
	head.Next = nil
	for cur != nil {
		loc := len(sortedList)
		for i := len(sortedList) - 1; i >= 0; i-- {
			if cur.Val < sortedList[i].Val {
				loc = i
			}
		}
		target := cur
		cur = cur.Next
		if loc > 0 {
			sortedList[loc-1].Next = target
		}
		if loc < len(sortedList) {
			target.Next = sortedList[loc]
		}
		tmp := append(make([]*ListNode, 0), sortedList[0:loc]...)
		tmp = append(tmp, []*ListNode{target}...)
		sortedList = append(tmp, sortedList[loc:]...)
		sortedList[len(sortedList)-1].Next = nil
	}
	return sortedList[0]
}

func build(list []int) *ListNode {
	head := &ListNode{list[0], nil}
	cur := head
	for i := 1; i < len(list); i++ {
		node := &ListNode{list[i], nil}
		cur.Next = node
		cur = node
	}
	return head
}

func print(head *ListNode) {
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	fmt.Println(result)
}

func main() {
	head := build([]int{6, 4, 5, 2, 1, 3})
	after := insertionSortList(head)
	print(after)
}
