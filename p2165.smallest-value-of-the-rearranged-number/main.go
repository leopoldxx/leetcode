package main

import (
	"sort"
)

func splitNum(num int64) (res []int, zero []int64) {
	for num > 0 {
		m := num % 10
		num = num / 10
		if m == 0 {
			zero = append(zero, m)
		} else {
			res = append(res, int(m))
		}
	}
	return
}

func cat(nums []int, zeros []int64, desc bool) (res int64) {
	if desc {
		for i := 0; i < len(nums); i++ {
			res = res*10 + int64(nums[i])
		}
		for i := 0; i < len(zeros); i++ {
			res = res * 10
		}
		return
	}
	if len(nums) > 0 {
		res = int64(nums[0])
		for i := 0; i < len(zeros); i++ {
			res = res * 10
		}
	}
	for i := 1; i < len(nums); i++ {
		res = res*10 + int64(nums[i])
	}
	return
}
func bigestNumber(num int64) int64 {
	nums, zeros := splitNum(num)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return cat(nums, zeros, true)
}

func smallestNumber(num int64) int64 {
	if num < 0 {
		return -bigestNumber(-num)
	}

	nums, zeros := splitNum(num)
	sort.Ints(nums)
	return cat(nums, zeros, false)
}
