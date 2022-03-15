package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var (
		pre     *ListNode
		newHead *ListNode
	)
	for head != nil {
		pre = newHead
		newHead = head
		head = head.Next
		newHead.Next = pre
	}
	return newHead
}
