package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func deepCopy(src *Node, into *Node, accessed map[*Node]*Node) {
	into.Val = src.Val
	accessed[src] = into
	if len(src.Neighbors) == 0 {
		return
	}

	into.Neighbors = make([]*Node, len(src.Neighbors))
	for i := range src.Neighbors {
		if _, exists := accessed[src.Neighbors[i]]; exists {
			into.Neighbors[i] = accessed[src.Neighbors[i]]
		} else {
			into.Neighbors[i] = &Node{Val: src.Neighbors[i].Val}
			deepCopy(src.Neighbors[i], into.Neighbors[i], accessed)
		}
	}
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	accessed := map[*Node]*Node{}
	copied := &Node{Val: node.Val}
	deepCopy(node, copied, accessed)

	return copied

}
