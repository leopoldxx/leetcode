package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapTwo(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	l1 := head
	if head.Next == nil {
		return head, nil
	}
	l2 := head.Next
	tail := head.Next.Next

	l2.Next = l1
	l1.Next = tail
	return l2, tail
}
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	newHead, next := swapTwo(head)
	pre := newHead.Next

	for next != nil {
		pre.Next, next = swapTwo(next)
		if pre.Next != nil {
			pre = pre.Next.Next
		}
	}
	return newHead
}
