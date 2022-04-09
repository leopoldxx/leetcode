package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	it := head
	var pre *ListNode
	for it != nil {
		if pre == nil {
			pre = it
			it = it.Next
		} else if pre.Val == it.Val {
			pre.Next = it.Next
			it.Next = nil
			it = pre.Next
		} else {
			pre = it
			it = it.Next
		}
	}
	return head
}
