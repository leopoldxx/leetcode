package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	listLen := 0
	it := head
	for it != nil {
		listLen++
		it = it.Next
	}
	if n == listLen {
		return head.Next
	}

	step := listLen - n

	it = head
	pre := head
	for step > 0 {
		pre = it
		it = it.Next
		step--
	}
	pre.Next = it.Next
	return head

}
