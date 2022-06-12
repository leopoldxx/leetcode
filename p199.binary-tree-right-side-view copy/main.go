package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	layer := []*TreeNode{root}
	res := []int{root.Val}
	for len(layer) > 0 {
		newLayer := []*TreeNode{}
		for i := 0; i < len(layer); i++ {
			if layer[i].Left != nil {
				newLayer = append(newLayer, layer[i].Left)
			}
			if layer[i].Right != nil {
				newLayer = append(newLayer, layer[i].Right)
			}
		}
		if len(newLayer) > 0 {
			res = append(res, newLayer[len(newLayer)-1].Val)
		}
		layer = newLayer
	}
	return res
}
