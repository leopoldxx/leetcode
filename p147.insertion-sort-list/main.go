package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre, it := head, head.Next
	for it != nil {
		pre.Next = it.Next
		it.Next = nil
		if it.Val < head.Val {
			it.Next = head
			head = it
			it = pre.Next
			continue
		}

		it2Pre := head
		it2 := head.Next
		for it2 != pre.Next {
			if it.Val < it2.Val {
				it.Next = it2
				it2Pre.Next = it
				it = nil
				break
			} else {
				it2Pre = it2
				it2 = it2.Next
			}
		}
		if it != nil {
			it.Next = pre.Next
			pre.Next = it
			pre = it
			it = it.Next
		} else {
			it = pre.Next
		}

	}
	return head
}
