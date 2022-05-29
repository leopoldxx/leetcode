package main

import (
	"testing"
)

func buildTree(preorder []int, inorder []int) *Node {
	if len(preorder) == 0 {
		return nil
	} else if len(preorder) == 1 {
		return &Node{
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
	return &Node{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:1+idx], inorder[0:idx]),
		Right: buildTree(preorder[1+idx:], inorder[idx+1:]),
	}
}

func TestSolution(t *testing.T) {
	tree := buildTree(
		[]int{1, 2, 4, 5, 3, 6, 7},
		[]int{4, 2, 5, 1, 6, 3, 7},
	)
	connect(tree)
}
