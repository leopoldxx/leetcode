package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getListLength(head *ListNode) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}
func swapNode(n1, n2 *ListNode) {
	n1.Val, n2.Val = n2.Val, n1.Val
}
func swapNodes(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	length := getListLength(head)
	node1Pos := k
	node2Pos := length - k + 1

	var node1, node2 *ListNode
	it := head
	itPos := 1
	for node1 == nil || node2 == nil {
		if itPos == node1Pos {
			node1 = it
		}
		if itPos == node2Pos {
			node2 = it
		}
		itPos++
		it = it.Next
	}
	swapNode(node1, node2)
	return head
}
