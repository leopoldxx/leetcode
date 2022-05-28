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
		preSum += root.Val
		path = append(path, root)
		dfs(root.Left, target, preSum, path, res)
		path = path[:len(path)-1]
		preSum -= root.Val
	}
	if root.Right != nil {
		preSum += root.Val
		path = append(path, root)
		dfs(root.Right, target, preSum, path, res)
		path = path[:len(path)-1]
		preSum -= root.Val
	}
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	dfs(root, targetSum, 0, []*TreeNode{}, &res)
	return res
}
