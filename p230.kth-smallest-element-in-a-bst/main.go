package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraverse(nodeSlice []*TreeNode, root *TreeNode) []*TreeNode {
	if root == nil {
		return nodeSlice
	}
	nodeSlice = inorderTraverse(nodeSlice, root.Left)
	nodeSlice = append(nodeSlice, root)
	nodeSlice = inorderTraverse(nodeSlice, root.Right)
	return nodeSlice
}
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	nodeSlice := make([]*TreeNode, 0, 10000)
	nodeSlice = inorderTraverse(nodeSlice, root)
	return nodeSlice[k-1].Val
}
