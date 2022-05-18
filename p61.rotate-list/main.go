package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func countAndTail(head *ListNode) (l int, tail *ListNode) {
	for head != nil {
		l++
		tail = head
		head = head.Next
	}
	return
}
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	n, tail := countAndTail(head)
	if n <= k {
		k %= n
	}
	if k == 0 {
		return head
	}

	step := n - k
	newHead := head
	preNewHead := newHead
	for s := 0; s < step; s++ {
		preNewHead = newHead
		newHead = newHead.Next
	}
	tail.Next = head
	preNewHead.Next = nil
	return newHead
}
