package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}

	for len(queue) > 0 {
		nextQueue := []*Node{}
		for i := 1; i < len(queue); i++ {
			queue[i-1].Next = queue[i]
			if queue[i-1].Left != nil {
				nextQueue = append(nextQueue, queue[i-1].Left)
			}
			if queue[i-1].Right != nil {
				nextQueue = append(nextQueue, queue[i-1].Right)
			}
		}
		if queue[len(queue)-1].Left != nil {
			nextQueue = append(nextQueue, queue[len(queue)-1].Left)
		}
		if queue[len(queue)-1].Right != nil {
			nextQueue = append(nextQueue, queue[len(queue)-1].Right)
		}
		queue = nextQueue
	}
	return root
}
