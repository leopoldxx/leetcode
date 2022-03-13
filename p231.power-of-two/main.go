package main

func countBitNum(n int) int {
	var count int
	for n > 0 {
		n = n & (n - 1)
		count++
	}
	return count
}
func isPowerOfTwo(n int) bool {
	return countBitNum(n) == 1
}
