package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		pre := head
		head = head.Next
		pre.Next = nil
	}
	if head == nil {
		return nil
	}
	pre, current := head, head.Next

	for current != nil {
		if current.Val == val {
			f := current

			pre.Next = current.Next
			current = current.Next

			f.Next = nil
			continue
		}

		pre = current
		current = current.Next
	}

	return head
}
