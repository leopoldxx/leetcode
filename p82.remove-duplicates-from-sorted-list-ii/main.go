package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	firstVal := head.Val
	for head != nil {
		if head != nil && head.Next != nil && head.Val == head.Next.Val {
			for head != nil && head.Val == firstVal {
				head = head.Next
			}
			if head != nil {
				firstVal = head.Val
			}
		} else {
			break
		}
	}

	if head == nil || head.Next == nil {
		return head
	}

	last := head
	it := head.Next
	preVal := it.Val
	for it != nil {
		if it.Next != nil && it.Val == it.Next.Val {
			for it != nil && it.Val == preVal {
				it = it.Next
			}
			if it != nil {
				preVal = it.Val
			}
		} else if it.Next != nil {
			last.Next = it
			last = it
			it = it.Next
			preVal = it.Val
		} else if last != it {
			last.Next = it
			last = it
			return head
		} else {
			return head
		}
	}
	if last != it {
		last.Next = it
		last = it
		return head
	}
	return head
}
