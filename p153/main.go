package main

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	left, right := 0, len(nums)-1
	for left <= right {
		if left == right {
			return nums[left]
		}
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return 0
}
