package main

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sepList(head *ListNode) (left, mid, right *ListNode) {
	fast, slow, slowPre := head, head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slowPre = slow
		slow = slow.Next
	}
	slowPre.Next = nil
	if head == slow {
		head = nil
	}
	return head, slow, slow.Next
}
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	left, mid, right := sepList(head)
	node := &TreeNode{
		Val: mid.Val,
	}
	if left != nil {
		node.Left = sortedListToBST(left)
	}
	if right != nil {
		node.Right = sortedListToBST(right)
	}
	return node
}
