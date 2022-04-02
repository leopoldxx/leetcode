package main

var isBadVersion func(version int) bool

func firstBadVersion(n int) int {
	left, right := 1, n
	for {
		if left > right {
			return left
		}
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

}
