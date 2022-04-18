package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) (newHead *ListNode, tail *ListNode) {
	tail = head
	var pre *ListNode

	for head != nil {
		pre = newHead
		newHead = head
		head = head.Next

		newHead.Next = pre
	}
	return newHead, tail

}
func equal(l, r *ListNode) bool {
	for {
		if l == r {
			return true
		}
		if l.Val != r.Val {
			return false
		}
		l = l.Next
		r = r.Next
	}

}
func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		if fast != nil {
			slow = slow.Next
		}
	}
	l2, t2 := reverse(slow.Next)

	if fast != nil {
		t2.Next = slow
	}
	slow.Next = nil

	return equal(head, l2)
}
