package main

func findLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for {
		if left >= right {
			return right
		}
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
}

func findRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for {
		if left >= right {
			return right
		}
		mid := left + (right-left)/2
		if right-left == 1 {
			mid += 1
		}
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := findLeft(nums, target)
	if nums[left] != target {
		return []int{-1, -1}
	}
	return []int{left, findRight(nums, target)}

}
