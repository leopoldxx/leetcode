package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	stack := []*Node{}
	origin := root
	for root.Next != nil || root.Child != nil || len(stack) > 0 {
		next := root.Next
		if root.Child != nil {
			if next != nil {
				stack = append(stack, next)
			}
			root.Next = root.Child
			root.Next.Prev = root
			root.Child = nil
		}
		if root.Next != nil {
			root = root.Next
			continue
		} else if len(stack) > 0 {
			next := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root.Next = next
			root.Next.Prev = root
			root.Child = nil
			root = next
			continue
		} else {
			return origin
		}
	}
	return origin
}
