package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// stop condition:
	// 1. the leaf;
	// 2. we found one of p or q;
	if root == nil || root == p || root == q {
		return root
	}

	// found in left subtree
	left := lowestCommonAncestor(root.Left, p, q)
	// found in right subtree
	right := lowestCommonAncestor(root.Right, p, q)

	// we found the two nodes in the subtrees, then root is the target node
	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}
