package main

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	major := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if major == nums[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			major = nums[i]
			count = 1
		}
	}
	count = 0
	for i := 0; i < len(nums); i++ {
		if major == nums[i] {
			count++
		}
	}
	if count > len(nums)/2 {
		return major
	}
	return 0
}
