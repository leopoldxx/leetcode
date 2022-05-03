package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	} else if root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}

	nodes := make([]*TreeNode, 0, 100)
	newnodes := make([]*TreeNode, 0, 100)
	res := make([][]int, 0, 100)

	nodes = append(nodes, root)
	res = append(res, []int{root.Val})

	lefttoright := true
	levelIndex := 0

	for levelIndex < len(res) {
		newres := []int{}
		newnodes = newnodes[:0]
		if lefttoright {
			for j := len(nodes) - 1; j >= 0; j-- {
				if nodes[j].Right != nil {
					newnodes = append(newnodes, nodes[j].Right)
					newres = append(newres, nodes[j].Right.Val)
				}
				if nodes[j].Left != nil {
					newnodes = append(newnodes, nodes[j].Left)
					newres = append(newres, nodes[j].Left.Val)
				}
			}
		} else {
			for j := len(nodes) - 1; j >= 0; j-- {
				if nodes[j].Left != nil {
					newnodes = append(newnodes, nodes[j].Left)
					newres = append(newres, nodes[j].Left.Val)
				}
				if nodes[j].Right != nil {
					newnodes = append(newnodes, nodes[j].Right)
					newres = append(newres, nodes[j].Right.Val)
				}
			}
		}
		if len(newres) == 0 {
			return res
		}
		nodes, newnodes = newnodes, nodes

		res = append(res, newres)
		lefttoright = !lefttoright
		levelIndex++
	}
	return res
}
