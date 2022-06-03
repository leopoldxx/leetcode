package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	} else if len(preorder) == 1 {
		return &TreeNode{
			Val: preorder[0],
		}
	}
	v := preorder[0]
	idx := -1
	for i, n := range inorder {
		if v == n {
			idx = i
			break
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:1+idx], inorder[0:idx]),
		Right: buildTree(preorder[1+idx:], inorder[idx+1:]),
	}
}

func TestSolution(t *testing.T) {
	tree := buildTree(
		[]int{7, 3, 15, 9, 20},
		[]int{3, 7, 9, 15, 20},
	)
	obj := Constructor(tree)
	_ = obj.Next()
	_ = obj.Next()
	_ = obj.Next()
	_ = obj.Next()
	_ = obj.Next()
	assert.False(t, obj.HasNext())
}
