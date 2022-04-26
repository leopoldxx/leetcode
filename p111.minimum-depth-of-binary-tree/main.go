package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := minDepth(root.Left)
	r := minDepth(root.Right)
	if l == 0 && r == 0 {
		return 1
	} else if l == 0 {
		return r + 1
	} else if r == 0 {
		return l + 1
	} else {
		return min(l, r) + 1
	}
}
