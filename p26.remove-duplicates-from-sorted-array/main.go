package main

func removeDuplicates(nums []int) int {
	left, right := 0, 1
	for {
		if right >= len(nums) {
			break
		}
		if nums[left] == nums[right] {
			right++
			continue
		}
		left++
		nums[left] = nums[right]
		right++
	}
	return left + 1
}
