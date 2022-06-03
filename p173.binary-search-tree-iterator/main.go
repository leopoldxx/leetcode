package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{
		stack: []*TreeNode{},
	}
	if root == nil {
		return it
	}
	it.stack = append(it.stack, root)
	left := root.Left
	for left != nil {
		it.stack = append(it.stack, left)
		left = left.Left
	}
	return it
}

func (this *BSTIterator) Next() int {
	if len(this.stack) == 0 {
		return -1
	}
	currentNode := this.stack[len(this.stack)-1] // top
	//this.stack[len(this.stack)-1] = nil
	this.stack = this.stack[:len(this.stack)-1]
	next := currentNode.Right
	for next != nil {
		this.stack = append(this.stack, next)
		next = next.Left
	}
	return currentNode.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}
