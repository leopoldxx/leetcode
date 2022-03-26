package main

type heap struct {
	data    []int
	compare func(i, j int) bool
}

func (h *heap) init(cmp func(i, j *int) bool) {
	h.compare = func(i, j int) bool {
		return cmp(&h.data[i], &h.data[j])
	}
}

func (h *heap) add(num int) {
	h.data = append(h.data, num)
	h.siftup(len(h.data) - 1)
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
func (h *heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
func (h *heap) len() int {
	return len(h.data)
}
func (h *heap) isLeaf(idx int) bool {
	return h.left(idx) >= h.len()
}
func (h *heap) lastIndex() int {
	return len(h.data) - 1
}
func (h *heap) inHeap(idx int) bool {
	return idx <= h.lastIndex()
}

func (h *heap) siftup(idx int) bool {
	changed := false
	for {
		if idx == 0 {
			break
		}
		p := h.parent(idx)
		if h.compare(p, idx) {
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
		if h.inHeap(right) && h.compare(right, left) {
			candidate = right
		}
		if h.compare(idx, candidate) {
			break
		}
		h.swap(idx, candidate)
		idx = candidate
		changed = true
	}

	return changed
}

func (h *heap) update(idx int, num int) {
	h.data[idx] = num
	if !h.siftdown(idx) {
		h.siftup(idx)
	}
}

func (h *heap) top() *int {
	if len(h.data) == 0 {
		return nil
	}
	return &h.data[0]
}

func (h *heap) pop() *int {
	if len(h.data) == 0 {
		return nil
	}
	v := h.data[0]
	lastIndex := h.lastIndex()
	h.swap(0, lastIndex)
	h.data = h.data[:lastIndex]
	h.siftdown(0)
	return &v
}

func (h *heap) build(input ...int) {
	for _, num := range input {
		h.add(num)
	}
}

type MedianFinder struct {
	maxHeap   *heap
	minHeap   *heap
	median    int
	medianSet bool
}

func Constructor() MedianFinder {
	maxHeap := &heap{}
	maxHeap.init(func(i, j *int) bool {
		return *j < *i
	})
	minHeap := &heap{}
	minHeap.init(func(i, j *int) bool {
		return *i < *j
	})
	return MedianFinder{
		maxHeap:   maxHeap,
		minHeap:   minHeap,
		median:    0,
		medianSet: false,
	}
}

func sort(l, r int) (min, max int) {
	if l < r {
		return l, r
	}
	return r, l
}

func (this *MedianFinder) AddNum(num int) {
	if this.medianSet {
		min, max := sort(num, this.median)
		this.medianSet = false
		this.maxHeap.add(min)
		this.minHeap.add(max)
		return
	}
	this.medianSet = true
	this.median = num
	l := this.maxHeap.top()
	if l == nil {
		return
	}
	r := this.minHeap.top()
	if num < *l {
		this.median = *l
		this.maxHeap.update(0, num)
	} else if num > *r {
		this.median = *r
		this.minHeap.update(0, num)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.medianSet {
		return float64(this.median)
	}
	return (float64(*this.maxHeap.top()) + float64(*this.minHeap.top())) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
