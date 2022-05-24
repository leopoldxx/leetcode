package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	head = dummy
	leftPointerPre := head
	leftPointer := head

	for i := 0; i < left; i++ {
		leftPointerPre = leftPointer
		leftPointer = leftPointer.Next
	}

	n := right - left

	rightPointerPre := leftPointer
	rightPointer := leftPointer.Next

	for i := 0; i < n; i++ {
		tmp := rightPointer
		rightPointer = rightPointer.Next
		tmp.Next = rightPointerPre
		rightPointerPre = tmp
	}

	leftPointerPre.Next = rightPointerPre
	leftPointer.Next = rightPointer

	return dummy.Next

}
