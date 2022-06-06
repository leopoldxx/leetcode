package main

func minSubArrayLen(target int, nums []int) int {
	left := 0
	right := 0
	sum := 0
	var minCount *int
	for right <= len(nums)-1 {
		sum += nums[right]
		for sum >= target {
			count := right - left + 1
			if minCount == nil {
				minCount = &count
			} else if *minCount > count {
				*minCount = count
			}
			sum -= nums[left]
			left++
		}
		right++
	}
	if minCount != nil {
		return *minCount
	}
	return 0
}
