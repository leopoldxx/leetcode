package main

type item struct {
	val     int
	idx     int
	heapIdx int
}

type maxheap struct {
	indice map[int]*item
	data   []*item
}

func (h *maxheap) init() {
	h.indice = make(map[int]*item)
}
func (h *maxheap) last() int {
	return len(h.data) - 1
}

func (h *maxheap) siftup(idx int) {
	for {
		p := (idx - 1) / 2
		if p == idx { // idx == 0
			return
		}
		if h.data[idx].val < h.data[p].val {
			return
		}
		h.data[idx], h.data[p] = h.data[p], h.data[idx]
		h.data[idx].heapIdx = idx
		h.data[p].heapIdx = p
		idx = p
	}
}
func (h *maxheap) siftdown(idx int) (changed bool) {
	for {
		left := idx*2 + 1
		if left >= len(h.data) {
			return
		}
		cand := left
		if right := left + 1; right < len(h.data) && h.data[right].val > h.data[left].val {
			cand = right
		}

		if h.data[cand].val <= h.data[idx].val {
			return
		}
		changed = true
		h.data[cand], h.data[idx] = h.data[idx], h.data[cand]
		h.data[cand].heapIdx = cand
		h.data[idx].heapIdx = idx
		idx = cand
	}
}
func (h *maxheap) push(v item) {
	h.indice[v.idx] = &v
	h.data = append(h.data, &v)
	v.heapIdx = h.last()
	h.siftup(h.last())
}

func (h *maxheap) top() *item {
	if len(h.data) == 0 {
		return nil
	}
	return h.data[0]
}

func (h *maxheap) remove(idx int) {
	it, ok := h.indice[idx]
	if !ok {
		return
	}

	heapIdx := it.heapIdx
	if heapIdx == h.last() {
		h.data[h.last()].heapIdx = -1
		h.data = h.data[:h.last()]
		return
	}

	h.data[heapIdx], h.data[h.last()] = h.data[h.last()], h.data[heapIdx]
	h.data[heapIdx].heapIdx = heapIdx
	h.data[h.last()].heapIdx = -1
	h.data = h.data[:h.last()]

	delete(h.indice, idx)

	h.siftdown(heapIdx)
	h.siftup(heapIdx)
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	} else if len(nums) < k {
		return nil
	}

	maxSumSlice := make([]int, 0, len(nums))
	heap := &maxheap{}
	heap.init()

	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			heap.push(item{val: nums[i], idx: i})
		} else {
			heap.push(item{val: nums[i], idx: i})
			maxSumSlice = append(maxSumSlice, heap.top().val)
			heap.remove(i - k + 1)
		}

	}
	return maxSumSlice
}
