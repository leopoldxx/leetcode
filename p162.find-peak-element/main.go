package main

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	n := len(nums)

	if nums[0] > nums[1] {
		return 0
	}
	if nums[n-1] > nums[n-2] {
		return n - 1
	}

	for i := 1; i < n-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}
	return -1
}
