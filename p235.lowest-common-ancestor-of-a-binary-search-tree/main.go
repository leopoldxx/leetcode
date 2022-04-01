package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	if p.Val > q.Val {
		p, q = q, p
	}
	it := root
	for {
		if (it.Val > p.Val && it.Val < q.Val) || (it == p) || (it == q) {
			return it
		} else if it.Val > p.Val && it.Val > q.Val {
			it = it.Left
		} else {
			it = it.Right
		}
	}
}
