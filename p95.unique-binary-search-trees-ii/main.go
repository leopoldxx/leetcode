package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildBST(ns []int) []*TreeNode {
	if len(ns) == 1 {
		return []*TreeNode{
			{
				Val: ns[0],
			},
		}
	}
	n := len(ns)
	result := []*TreeNode{}
	for i := 0; i < n; i++ {
		if i == 0 {
			subres := buildBST(ns[1:])
			for j := 0; j < len(subres); j++ {
				result = append(result, &TreeNode{
					Val:   ns[i],
					Right: subres[j],
				})
			}
		} else if i == n-1 {
			subres := buildBST(ns[:n-1])
			for j := 0; j < len(subres); j++ {
				result = append(result, &TreeNode{
					Val:  ns[i],
					Left: subres[j],
				})
			}
		} else {
			subresleft := buildBST(ns[:i])
			subresright := buildBST(ns[i+1:])
			for ii := 0; ii < len(subresleft); ii++ {
				for jj := 0; jj < len(subresright); jj++ {
					result = append(result, &TreeNode{
						Val:   ns[i],
						Left:  subresleft[ii],
						Right: subresright[jj],
					})
				}
			}
		}
	}

	return result
}

func generateTrees(n int) []*TreeNode {
	ns := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		ns = append(ns, i)
	}
	return buildBST(ns)
}
