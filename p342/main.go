package main

func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && (n == n&0x555555555555555)
}
