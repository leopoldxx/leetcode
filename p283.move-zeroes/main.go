package main

func moveZeroes(nums []int) {
	idx := 0
	target := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			continue
		}
		nums[idx] = nums[i]
		idx++
	}
	for i := idx; i < len(nums); i++ {
		nums[i] = target
	}
}
