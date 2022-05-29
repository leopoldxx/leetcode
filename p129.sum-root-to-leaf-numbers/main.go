package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, path []*TreeNode, res *int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		var val int
		for i := 0; i < len(path); i++ {
			val = val*10 + path[i].Val
		}
		val = val*10 + root.Val
		*res += val
		return
	}
	if root.Left != nil {
		dfs(root.Left, append(path, root), res)
	}
	if root.Right != nil {
		dfs(root.Right, append(path, root), res)
	}
}

func sumNumbers(root *TreeNode) int {
	var res int
	dfs(root, []*TreeNode{}, &res)
	return res
}
