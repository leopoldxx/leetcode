package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func listLen(head *ListNode) int {
	next := head
	count := 0
	for next != nil {
		count++
		next = next.Next
	}
	return count
}
func middleNode(head *ListNode) *ListNode {
	length := listLen(head)
	if length == 1 {
		return head
	}
	half := (length + 1) / 2
	first := head
	second := head
	for half > 0 {
		first = first.Next
		half--
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	return second
}
