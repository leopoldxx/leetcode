package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	if len(nums) == 0 || k <= 1 {
		return 0
	}
	var (
		result  = 0
		product = 1
		left    = 0
	)
	for right := 0; right < len(nums); right++ {
		product *= nums[right]
		for product >= k {
			product /= nums[left]
			left++
		}
		result = result + (right - left + 1)
	}
	return result
}
