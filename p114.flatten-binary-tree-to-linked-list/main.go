package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLast(root *TreeNode) (last *TreeNode) {
	if root == nil {
		return nil
	}
	if root.Right == nil && root.Left == nil {
		return root
	} else if root.Right != nil {
		return findLast(root.Right)
	} else {
		return findLast(root.Left)
	}
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		last := findLast(root.Left)
		last.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
	flatten(root.Right)
}
