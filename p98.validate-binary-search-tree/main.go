package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidWithMinMax(root *TreeNode) (bool, *int, *int) {
	if root == nil {
		return true, nil, nil
	} else if root.Left == nil && root.Right == nil {
		return true, &root.Val, &root.Val
	} else if root.Left != nil && root.Left.Val >= root.Val {
		return false, nil, nil
	} else if root.Right != nil && root.Right.Val <= root.Val {
		return false, nil, nil
	}

	left, lmin, lmax := isValidWithMinMax(root.Left)
	if !left || (lmax != nil && *lmax >= root.Val) {
		return false, nil, nil
	}
	right, rmin, rmax := isValidWithMinMax(root.Right)
	if !right || (rmin != nil && *rmin <= root.Val) {
		return false, nil, nil
	}
	if lmin == nil {
		lmin = &root.Val
	}
	if rmax == nil {
		rmax = &root.Val
	}

	return true, lmin, rmax
}

func isValidBST(root *TreeNode) bool {
	res, _, _ := isValidWithMinMax(root)

	return res

}
