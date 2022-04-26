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

type maxDepth func(root *TreeNode) int

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var maxDiameter int
	var depth maxDepth
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		ld := depth(root.Left)
		rd := depth(root.Right)
		if maxDiameter < ld+rd {
			maxDiameter = ld + rd
		}
		return 1 + max(ld, rd)
	}
	depth(root)
	return maxDiameter
}
