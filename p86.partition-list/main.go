package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var (
		h1, n1 *ListNode
		h2, n2 *ListNode
	)
	for head != nil {
		if head.Val < x {
			if h1 == nil {
				h1, n1 = head, head
			} else {
				n1.Next = head
				n1 = n1.Next
			}
			head = head.Next
			n1.Next = nil
		} else {
			if h2 == nil {
				h2, n2 = head, head
			} else {
				n2.Next = head
				n2 = n2.Next
			}
			head = head.Next
			n2.Next = nil
		}
	}
	if n1 == nil {
		return h2
	}
	n1.Next = h2
	return h1
}
