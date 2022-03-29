package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSame(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} else if root != nil && subRoot == nil {
		return false
	} else if root == nil && subRoot != nil {
		return false
	} else if root.Val != subRoot.Val {
		return false
	}
	return isSame(root.Left, subRoot.Left) && isSame(root.Right, subRoot.Right)
}
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} else if root != nil && subRoot == nil {
		return false
	} else if root == nil && subRoot != nil {
		return false
	}
	return isSame(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
