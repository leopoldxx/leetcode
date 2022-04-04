package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTargetWithCache(root *TreeNode, k int, cache map[int]struct{}) bool {
	if root == nil {
		return false
	}
	if _, exists := cache[k-root.Val]; exists {
		return true
	}
	cache[root.Val] = struct{}{}

	return findTargetWithCache(root.Left, k, cache) || findTargetWithCache(root.Right, k, cache)
}

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	cache := map[int]struct{}{}
	cache[root.Val] = struct{}{}
	return findTargetWithCache(root.Left, k, cache) || findTargetWithCache(root.Right, k, cache)
}
