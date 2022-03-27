package main

func missingNumber(nums []int) int {
	diff := 0
	for i := 0; i < len(nums); i++ {
		diff += (nums[i] - i)
	}
	return len(nums) - diff
}
