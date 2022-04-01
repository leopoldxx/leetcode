package main

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for {
		if left > right {
			return left
		}
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}
