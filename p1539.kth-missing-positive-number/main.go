package main

func findKthPositive(arr []int, k int) int {
	left, right := 0, len(arr)-1
	mid, missed := 0, 0
	for {
		if left > right {
			break
		}
		mid = left + (right-left)/2
		missed = arr[mid] - mid - 1
		if missed >= k {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right + k + 1
}
