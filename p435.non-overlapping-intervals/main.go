package main

import "sort"

type intervalList [][]int

func (il intervalList) Len() int {
	return len(il)
}
func (il intervalList) Less(i, j int) bool {
	return il[i][1] < il[j][1]
}
func (il intervalList) Swap(i, j int) {
	il[i], il[j] = il[j], il[i]
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 1 {
		return 0
	}
	sort.Sort(intervalList(intervals))
	count := 1
	lastEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= lastEnd {
			lastEnd = intervals[i][1]
			count++
		}
	}
	return len(intervals) - count

}
