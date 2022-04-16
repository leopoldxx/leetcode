package main

import (
	"sort"
)

func findPeakFromEnd(nums []int) int {
	for j := len(nums) - 1; j > 0; j-- {
		if nums[j] > nums[j-1] {
			return j
		}
	}
	return 0
}

func swap(nums []int, pre int, peak int) {
	target := peak
	for j := len(nums) - 1; j >= peak; j-- {
		if nums[j] > nums[pre] {
			target = j
			break
		}
	}

	nums[target], nums[pre] = nums[pre], nums[target]
}

func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	peak := findPeakFromEnd(nums)
	if peak == 0 {
		sort.Ints(nums)
		return
	}
	swap(nums, peak-1, peak)
	sort.Ints(nums[peak:])
}
