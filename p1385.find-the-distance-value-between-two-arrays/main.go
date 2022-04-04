package main

import "sort"

func findNumPosInArray(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
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

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	count := 0
	for _, num := range arr1 {
		pos := findNumPosInArray(arr2, num)
		if pos == 0 {
			dd := num - arr2[pos]
			if dd <= d && dd >= -d {
				count++
			}
		} else if pos == len(arr2) {
			dd := num - arr2[pos-1]
			if dd <= d && dd >= -d {
				count++
			}
		} else {
			dd1 := num - arr2[pos]
			dd2 := num - arr2[pos-1]
			if (dd1 <= d && dd1 >= -d) || (dd2 <= d && dd2 >= -d) {
				count++
			}
		}
	}
	return len(arr1) - count
}
