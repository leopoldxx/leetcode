package main

func singleNonDuplicate(nums []int) int {
	for len(nums) > 1 {
		idx := len(nums) / 2
		if nums[idx] == nums[idx-1] {
			if (idx+1)%2 == 0 {
				nums = nums[idx+1:]
			} else {
				nums = nums[:idx+1]
			}
		} else if nums[idx] == nums[idx+1] {
			if (len(nums)-idx)%2 == 0 {
				nums = nums[:idx]
			} else {
				nums = nums[idx:]
			}
		} else {
			return nums[idx]
		}
	}
	return nums[0]
}
