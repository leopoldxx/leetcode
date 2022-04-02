package main

var guess func(num int) int

func guessNumber(n int) int {
	left, right := 1, n
	for {
		mid := left + (right-left)/2
		g := guess(mid)
		if g == 0 {
			return mid
		} else if g == 1 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
