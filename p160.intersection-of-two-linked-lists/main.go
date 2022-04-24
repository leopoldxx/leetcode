package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	itA := headA
	itB := headB

	for itA != itB {
		if itA != nil {
			itA = itA.Next
		} else {
			itA = headB
		}
		if itB != nil {
			itB = itB.Next
		} else {
			itB = headA
		}
	}
	return itA
}
