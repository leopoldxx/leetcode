package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type cacheType struct {
	node  *TreeNode
	level int
	next  *cacheType
}
type queue struct {
	head *cacheType
	tail *cacheType
}

func (q *queue) push(node *cacheType) {
	if q.tail == nil {
		q.head, q.tail = node, node
		return
	}
	q.tail.next = node
	q.tail = node
}
func (q *queue) pop() *cacheType {
	if q.head == nil {
		return nil
	}
	n := q.head
	if q.head == q.tail {
		q.head, q.tail = nil, nil
		return n
	}
	q.head = q.head.next
	return n
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := &queue{}
	q.push(&cacheType{
		node:  root,
		level: 0,
	})
	cache := map[int][]int{}

	for {
		node := q.pop()
		if node == nil {
			break
		}
		cache[node.level] = append(cache[node.level], node.node.Val)
		if node.node.Left != nil {
			q.push(&cacheType{
				node:  node.node.Left,
				level: node.level + 1,
			})
		}
		if node.node.Right != nil {
			q.push(&cacheType{
				node:  node.node.Right,
				level: node.level + 1,
			})
		}
	}
	result := [][]int{}
	length := len(cache)
	for i := 0; i < length; i++ {
		result = append(result, cache[i])
	}
	return result

}
