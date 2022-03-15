package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(left *ListNode, right *ListNode) *ListNode {
	var head, next *ListNode
	for left != nil && right != nil {
		if left.Val < right.Val {
			if head == nil {
				head = left
				next = left
			} else {
				next.Next = left
				next = left
			}
			left = left.Next
		} else {
			if head == nil {
				head = right
				next = right
			} else {
				next.Next = right
				next = right
			}
			right = right.Next
		}
	}
	if left != nil {
		if next != nil {
			next.Next = left
		} else {
			return left
		}
	}
	if right != nil {
		if next != nil {
			next.Next = right
		} else {
			return right
		}
	}
	return head
}

func mergeSortList(head *ListNode, listLen int) *ListNode {
	if listLen == 0 {
		return nil
	}
	if listLen == 1 {
		head.Next = nil
		return head
	}
	leftHalfLen := listLen >> 1

	rightHalf := head
	for i := 0; i < leftHalfLen; i++ {
		rightHalf = rightHalf.Next
	}

	left := mergeSortList(head, leftHalfLen)
	right := mergeSortList(rightHalf, listLen-leftHalfLen)
	return merge(left, right)
}
func sortList(head *ListNode) *ListNode {
	listLen := 0
	cur := head
	for cur != nil {
		listLen++
		cur = cur.Next
	}
	return mergeSortList(head, listLen)
}
