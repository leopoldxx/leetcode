package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast || slow == fast.Next {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}
