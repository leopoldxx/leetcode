package main

type itemq struct {
	val  int
	pre  *itemq
	next *itemq
}

type deque struct {
	head *itemq
	tail *itemq
}

func (d *deque) init() {
}

func (d *deque) pushfront(v int) {
	item := itemq{
		val: v,
		pre: d.head,
	}
	if d.head == nil {
		d.head = &item
		d.tail = &item
		return
	}
	d.head.next = &item
	d.head = &item
}

func (d *deque) popfront() *itemq {
	if d.head == nil {
		return nil
	} else if d.head == d.tail {
		item := d.head
		d.head, d.tail = nil, nil
		return item
	}
	item := d.head
	d.head = d.head.pre

	item.pre = nil
	d.head.next = nil
	return item
}
func (d *deque) popback() *itemq {
	if d.tail == nil {
		return nil
	} else if d.head == d.tail {
		item := d.tail
		d.head, d.tail = nil, nil
		return item
	}
	item := d.tail
	d.tail = d.tail.next

	item.next = nil
	d.tail.pre = nil
	return item
}
func (d *deque) front() *itemq {
	if d.head == nil {
		return nil
	}
	return d.head
}
func (d *deque) back() *itemq {
	if d.tail == nil {
		return nil
	}
	return d.tail
}

func maxSlidingWindowDeque(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	} else if len(nums) < k {
		return nil
	}

	maxSumSlice := make([]int, 0, len(nums))

	q := &deque{}
	q.init()
	n := len(nums)
	for i := 0; i < n; i++ {
		if i < k-1 {
			c := nums[i]
			for item := q.front(); item != nil && nums[item.val] < c; {
				q.popfront()
				item = q.front()
			}
			q.pushfront(i)
		} else {
			c := nums[i]
			for item := q.front(); item != nil && nums[item.val] < c; {
				q.popfront()
				item = q.front()
			}
			q.pushfront(i)
			maxSumSlice = append(maxSumSlice, nums[q.back().val])
			for item := q.back(); item != nil && item.val <= i-k+1; {
				q.popback()
				item = q.back()
			}
		}
	}

	return maxSumSlice
}
