package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for {
		if left > right {
			return -1
		}
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
