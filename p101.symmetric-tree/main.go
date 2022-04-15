package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func equal(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	return equal(left.Left, right.Right) && equal(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return equal(root.Left, root.Right)
}
