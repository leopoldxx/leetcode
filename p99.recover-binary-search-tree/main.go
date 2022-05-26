package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(root *TreeNode, vals *[]int, nodes *[]*TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inorder(root.Left, vals, nodes)
	}
	*vals = append(*vals, root.Val)
	*nodes = append(*nodes, root)
	if root.Right != nil {
		inorder(root.Right, vals, nodes)
	}
}

func recoverTree(root *TreeNode) {
	vals := []int{}
	nodes := []*TreeNode{}
	inorder(root, &vals, &nodes)

	sort.Ints(vals)
	for i := 0; i < len(nodes); i++ {
		nodes[i].Val = vals[i]
	}

}
