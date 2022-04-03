package main

// idx is guaranteed that it's not on the array edge
func isPeak(arr []int, idx int) int {
	if arr[idx] > arr[idx-1] && arr[idx] > arr[idx+1] {
		return 0
	} else if arr[idx] > arr[idx-1] && arr[idx] < arr[idx+1] {
		return 1
	} else {
		return -1
	}
}

func peakIndexInMountainArray(arr []int) int {
	left, right := 1, len(arr)-2
	for {
		mid := left + (right-left)/2
		if res := isPeak(arr, mid); res == 0 {
			return mid
		} else if res == 1 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
