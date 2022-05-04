package main

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	i := 0
	j := len(nums) - 1

	count := 0
	for i < j {
		if nums[i]+nums[j] == k {
			count++
			i++
			j--
		} else if nums[i]+nums[j] < k {
			i++
		} else {
			j--
		}
	}
	return count
}
