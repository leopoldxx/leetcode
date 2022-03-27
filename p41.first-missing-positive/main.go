package main

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		if nums[i] < 0 || nums[i] > n {
			nums[i] = 0
		}
	}
	extraNPlus1 := 0
	virtualLen := n + 1
	for i := range nums {
		if nums[i] <= 0 {
			continue
		}
		mod := nums[i] % virtualLen
		if mod == n {
			extraNPlus1 += virtualLen
			continue
		}
		nums[mod] += virtualLen
	}
	for i := 1; i < n; i++ {
		if nums[i] < virtualLen {
			return i
		}
	}
	if extraNPlus1 == 0 {
		return n
	}
	return n + 1
}
