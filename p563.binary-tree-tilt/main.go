package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func dfs(root *TreeNode) (sum int, tilt int) {
	if root == nil {
		return 0, 0
	}
	leftSum, leftTilt := dfs(root.Left)
	rightSum, rightTilt := dfs(root.Right)
	sum = root.Val + leftSum + rightSum
	tilt = abs(leftSum-rightSum) + leftTilt + rightTilt
	return
}

func findTilt(root *TreeNode) int {
	_, tilt := dfs(root)
	return tilt
}
