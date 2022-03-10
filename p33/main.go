package main

func findMinIndex(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if left == right {
			return left
		}
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return 0
}
func binarySearch(nums []int, left, right, target int) int {
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if target == nums[0] {
			return 0
		} else {
			return -1
		}
	}
	minIndex := findMinIndex(nums)
	if target == nums[minIndex] {
		return minIndex
	} else if minIndex > 0 && nums[0] <= target {
		return binarySearch(nums, 0, minIndex-1, target)
	} else if minIndex > 0 && nums[0] > target {
		return binarySearch(nums, minIndex, len(nums)-1, target)
	} else {
		return binarySearch(nums, 0, len(nums)-1, target)
	}
}
