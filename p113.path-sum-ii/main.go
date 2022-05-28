package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, target int, preSum int, path []*TreeNode, res *[][]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if preSum+root.Val == target {
			pathVal := []int{}
			for i := 0; i < len(path); i++ {
				pathVal = append(pathVal, path[i].Val)
			}
			pathVal = append(pathVal, root.Val)
			*res = append(*res, pathVal)
			return
		}
		return
	}
	if root.Left != nil {
		dfs(root.Left, target, preSum+root.Val, append(path, root), res)
	}
	if root.Right != nil {
		dfs(root.Right, target, preSum+root.Val, append(path, root), res)
	}
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	dfs(root, targetSum, 0, []*TreeNode{}, &res)
	return res
}
