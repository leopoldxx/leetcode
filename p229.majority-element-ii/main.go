package main

import "sort"

func majorityElement(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	waterline := n / 3
	count := 0

	res := []int{}
	cursor := nums[0]
	count++
	for i := 1; i < len(nums); i++ {
		if nums[i] == cursor {
			count++
		} else {
			if count > waterline {
				res = append(res, cursor)
			}
			cursor = nums[i]
			count = 1
		}
	}
	if count > waterline {
		res = append(res, cursor)
	}
	return res
}
