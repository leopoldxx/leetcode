package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, path []*TreeNode, res *[]string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		var str string
		for i := 0; i < len(path); i++ {
			str += fmt.Sprintf("%d->", path[i].Val)
		}
		*res = append(*res, str+fmt.Sprintf("%d", root.Val))
		return
	}
	if root.Left != nil {
		dfs(root.Left, append(path, root), res)
	}
	if root.Right != nil {
		dfs(root.Right, append(path, root), res)
	}
}

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	dfs(root, []*TreeNode{}, &res)
	return res
}
