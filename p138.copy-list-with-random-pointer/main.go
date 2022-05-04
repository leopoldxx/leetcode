package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	mapper := map[*Node]*Node{}

	newNext := &Node{Val: head.Val}
	newHead := newNext
	mapper[head] = newNext

	next := head.Next
	for next != nil {
		newNode := &Node{Val: next.Val}
		mapper[next] = newNode
		newNext.Next = newNode
		newNext = newNode
		next = next.Next
	}

	newNext = newHead
	next = head
	for next != nil {
		if next.Random != nil {
			newNext.Random = mapper[next.Random]
		}
		next = next.Next
		newNext = newNext.Next
	}
	return newHead
}
