package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isBalancedWithHeight(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	} else if root.Left == nil && root.Right == nil {
		return true, 1
	}

	bl, hl := isBalancedWithHeight(root.Left)
	br, hr := isBalancedWithHeight(root.Right)

	mm := max(hl, hr)

	if !bl || !br {
		return false, mm + 1
	}

	mn := min(hl, hr)
	if mm-mn > 1 {
		return false, mm + 1
	}

	return true, mm + 1
}

func isBalanced(root *TreeNode) bool {
	b, _ := isBalancedWithHeight(root)
	return b
}
