package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	} else if n == 1 {
		return &TreeNode{
			Val: postorder[0],
		}
	}
	root := postorder[n-1]
	var index int
	for i := 0; i < n; i++ {
		if inorder[i] == root {
			index = i
			break
		}
	}
	return &TreeNode{
		Val:   root,
		Left:  buildTree(inorder[:index], postorder[:index]),
		Right: buildTree(inorder[index+1:], postorder[index:n-1]),
	}
}
