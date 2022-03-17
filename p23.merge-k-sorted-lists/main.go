package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type HeapList struct {
	data []*ListNode
}

func initHeapList(lists []*ListNode) *HeapList {
	h := &HeapList{
		data: lists,
	}
	h.buildHeap()
	return h
}

func (heap *HeapList) left(index int) int {
	return 2*index + 1
}

func (heap *HeapList) right(index int) int {
	return 2*index + 2
}

func (heap *HeapList) len() int {
	return len(heap.data)
}

func (heap *HeapList) buildHeap() {
	for current := heap.len()/2 - 1; current >= 0; current-- {
		heap.shiftDown(current)
	}
}

func (heap *HeapList) parent(index int) int {
	return (index - 1) >> 1
}

func (heap *HeapList) less(i, j int) bool {
	if heap.data[i] == nil {
		return true
	}
	if heap.data[j] == nil {
		return false
	}
	return heap.data[i].Val <= heap.data[j].Val
}

func (heap *HeapList) swap(i, j int) {
	heap.data[i], heap.data[j] = heap.data[j], heap.data[i]
}

func (heap *HeapList) isLeaf(i int) bool {
	return heap.left(i) >= heap.len()
}

func (heap *HeapList) shiftDown(index int) (shift bool) {
	for {
		if heap.isLeaf(index) {
			break
		}
		left := heap.left(index)
		candidate := left

		right := heap.right(index)
		if right < heap.len() && heap.less(right, left) {
			candidate = right
		}
		if heap.less(index, candidate) {
			break
		}
		heap.swap(index, candidate)
		shift = true
		index = candidate
	}
	return
}

func (heap *HeapList) shiftUp(index int) (shift bool) {
	for {
		p := heap.parent(index)
		if p == index {
			break
		}
		if heap.less(p, index) {
			break
		}
		heap.swap(p, index)
		shift = true
		index = p
	}
	return
}
func (heap *HeapList) pop() *ListNode {
	if heap.len() == 0 {
		return nil
	}
	elem := heap.data[0]
	for elem == nil {
		idx := heap.len() - 1
		heap.swap(0, idx)
		heap.data = heap.data[:idx]
		if heap.len() == 0 {
			return nil
		}
		heap.shiftDown(0)
		elem = heap.data[0]
	}
	if elem.Next == nil {
		idx := heap.len() - 1
		heap.swap(0, idx)
		heap.data = heap.data[:idx]
		heap.shiftDown(0)
		return elem
	}

	heap.data[0] = elem.Next
	elem.Next = nil
	heap.shiftDown(0)
	return elem
}

func mergeKLists(lists []*ListNode) *ListNode {
	var (
		head, next *ListNode
	)
	h := initHeapList(lists)
	for node := h.pop(); node != nil; node = h.pop() {
		if head == nil {
			head, next = node, node
			continue
		}
		next.Next = node
		next = next.Next
	}
	return head
}
