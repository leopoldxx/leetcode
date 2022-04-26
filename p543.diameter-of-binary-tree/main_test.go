package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	diameterOfBinaryTree(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	})
}
