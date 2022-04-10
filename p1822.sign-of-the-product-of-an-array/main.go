package main

func arraySign(nums []int) int {
	negative1Count := 0
	zeroCount := 0
	positive1Count := 0

	for _, num := range nums {
		switch {
		case num == 0:
			zeroCount++
		case num < 0:
			negative1Count++
		case num > 0:
			positive1Count++
		}
	}
	if zeroCount > 0 {
		return 0
	} else if negative1Count%2 == 1 {
		return -1
	}
	return 1
}
