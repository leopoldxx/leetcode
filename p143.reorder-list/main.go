package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre, current *ListNode = nil, head
	for current != nil {
		head = head.Next
		current.Next, pre, current = pre, current, head
	}
	return pre
}
func reorderList(head *ListNode) {
	slow, fast := head, head
	if fast == nil || fast.Next == nil {
		return
	}
	fast = fast.Next

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast.Next == nil {
			break
		}
		fast = fast.Next
	}
	rightHalf := reverseList(slow.Next)
	slow.Next = nil
	currentLeft, currentLeftNext := head, head.Next
	currentRight, currentRightNext := rightHalf, rightHalf.Next
	for currentLeft != nil {
		currentLeft.Next = currentRight
		currentRight.Next = currentLeftNext
		if currentLeftNext == nil {
			break
		}
		currentLeft, currentLeftNext = currentLeftNext, currentLeftNext.Next
		if currentRightNext == nil {
			break
		}
		currentRight, currentRightNext = currentRightNext, currentRightNext.Next
	}

}
