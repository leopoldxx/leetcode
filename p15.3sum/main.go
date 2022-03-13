package main

import "sort"

func binarySearch(nums []int, left, right, target int) int {
	for left <= right {
		mid := left + (right-left)/2

		if target == nums[mid] {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	type triplet [3]int
	filter := map[triplet]struct{}{}

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			complement := -(nums[i] + nums[j])
			k := binarySearch(nums, j+1, len(nums)-1, complement)
			if k == -1 {
				continue
			}
			filter[triplet{nums[i], nums[j], nums[k]}] = struct{}{}
		}
	}

	var result [][]int
	for t := range filter {
		result = append(result, []int{t[0], t[1], t[2]})
	}

	return result
}
