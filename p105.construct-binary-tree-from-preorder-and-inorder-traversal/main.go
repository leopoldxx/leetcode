package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
