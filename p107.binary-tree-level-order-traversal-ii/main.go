package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	res := [][]int{}

	for len(queue) > 0 {
		nextQueue := []*TreeNode{}
		nextLevelVals := []int{}
		for _, node := range queue {
			nextLevelVals = append(nextLevelVals, node.Val)
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		queue = nextQueue
		res = append([][]int{nextLevelVals}, res...)
	}
	return res
}
