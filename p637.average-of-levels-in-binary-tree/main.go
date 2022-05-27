package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func average(vals []int) float64 {
	if len(vals) == 0 {
		return 0
	}
	sum := 0
	for _, v := range vals {
		sum += v
	}
	return float64(sum) / float64(len(vals))
}
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	res := []float64{}

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
		res = append(res, average(nextLevelVals))
	}
	return res
}
