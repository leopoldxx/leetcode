package main

func search(nums []int, target int) bool {
	if len(nums) == 1 {
		return nums[0] == target
	}
	n := len(nums)
	left := 0
	right := n - 1

	for left <= right {
		if left == right {
			return nums[left] == target
		}
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] > nums[left] && (nums[left] < target && target < nums[mid]) {
			right = mid - 1
		} else if nums[mid] > nums[left] && (target > nums[mid] || target < nums[left]) {
			left = mid + 1
		} else if nums[mid] < nums[left] && (nums[mid] < target && target < nums[right]) {
			left = mid + 1
		} else if nums[mid] < nums[left] && (target < nums[mid] || target > nums[right]) {
			right = mid - 1
		} else {
			return search(nums[left:mid], target) || search(nums[mid+1:], target)
		}
	}
	return false
}
