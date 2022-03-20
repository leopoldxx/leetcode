package main

import "sort"

type stack struct {
	data [][]int
}

func (s *stack) push(e []int) *stack {
	s.data = append(s.data, e)
	return s
}
func (s *stack) len() int {
	return len(s.data)
}

func (s *stack) pop() []int {
	if len(s.data) == 0 {
		return nil
	}
	t := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return t
}
func (s *stack) raw() [][]int {
	return s.data
}

func mergeIntervals(left, right []int) ([]int, []int) {
	if left[0] <= right[0] && left[1] >= right[1] {
		return left, nil
	} else if right[0] <= left[0] && right[1] >= left[1] {
		return right, nil
	} else if left[0] <= right[0] && left[1] >= right[0] && left[1] <= right[1] {
		return []int{left[0], right[1]}, nil
	} else if right[0] <= left[0] && right[1] >= left[0] && right[1] <= left[1] {
		return []int{right[0], left[1]}, nil
	} else if left[1] < right[0] {
		return left, right
	} else {
		return right, left
	}
}

type intervalsType [][]int

func (it intervalsType) Less(i, j int) bool {
	return it[i][0] < it[j][0]
}
func (it intervalsType) Len() int {
	return len(it)
}
func (it intervalsType) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	sort.Sort(intervalsType(intervals))

	cache := &stack{}
	left := intervals[0]
	for i := 1; i < len(intervals); i++ {
		l, r := mergeIntervals(left, intervals[i])
		if r == nil {
			left = l
			continue
		}
		cache.push(l)
		left = r
	}
	return cache.push(left).raw()
}
