package main

type elem struct {
	num  int
	freq int
}
type heap struct {
	data     []elem
	capacity int
}

func (h *heap) init(capacity int) {
	h.capacity = capacity
}

func (h *heap) len() int {
	return len(h.data)
}
func (h *heap) add(e elem) {
	if h.len() < h.capacity {
		h.data = append(h.data, e)
		h.siftup(len(h.data) - 1)
		return
	}
	top := h.top()
	if top.freq > e.freq {
		return
	}
	h.data[0] = e
	h.siftdown(0)
}
func (h *heap) top() *elem {
	if len(h.data) == 0 {
		return nil
	}
	return &h.data[0]
}

func (h *heap) parent(idx int) int {
	return (idx - 1) >> 1
}
func (h *heap) left(idx int) int {
	return idx<<1 + 1
}
func (h *heap) right(idx int) int {
	return idx<<1 + 2
}
func (h *heap) isLeaf(idx int) bool {
	return h.left(idx) >= len(h.data)
}
func (h *heap) inheap(idx int) bool {
	return idx < h.len()
}
func (h *heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
func (h *heap) less(i, j int) bool {
	return h.data[i].freq < h.data[j].freq
}
func (h *heap) siftup(idx int) bool {
	changed := false
	for {
		if idx == 0 {
			break
		}
		p := h.parent(idx)
		if h.less(p, idx) {
			break
		}
		h.swap(p, idx)
		changed = true
		idx = p
	}
	return changed
}

func (h *heap) siftdown(idx int) bool {
	changed := false
	for {
		if h.isLeaf(idx) {
			break
		}
		left := h.left(idx)
		candidate := left
		right := h.right(idx)
		if h.inheap(right) && h.less(right, left) {
			candidate = right
		}
		if h.less(idx, candidate) {
			break
		}
		h.swap(idx, candidate)
		changed = true
		idx = candidate
	}
	return changed
}

func topKFrequent(nums []int, k int) []int {
	cache := map[int]int{}
	for _, num := range nums {
		cache[num]++
	}
	h := &heap{}
	h.init(k)
	for k, v := range cache {
		h.add(elem{num: k, freq: v})
	}
	result := []int{}
	for i := range h.data {
		result = append(result, h.data[i].num)
	}
	return result
}
