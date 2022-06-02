package main

func removeElement(nums []int, val int) int {
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[idx] = nums[i]
		idx++
	}
	return idx
}
