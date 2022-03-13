package main

import (
	"math"
)

const (
	MaxInt = int((^uint(0)) >> 1)
)

func maxPowerThree() int {
	max := 0
	num := uint(math.Pow(3, float64(20)))
	for {
		num *= 3
		if num > uint(MaxInt) {
			return max
		}
		max = int(num)
	}
}

func isPowerOfThree(n int) bool {
	return maxPowerThree()%n == 0
}

/*
func countBitNumBaseThree(n int) (int, int, int) {
	count1 := 0
	count2 := 0
	for n > 0 {
		mod := n % 3
		if mod == 1 {
			count1++
		} else if mod == 2 {
			count2++
		}
		n /= 3
	}
	return count1 + count2, count1, count2
}
func isPowerOfThree(n int) bool {
	count, count1, _ := countBitNumBaseThree(n)
	return count == 1 && count1 == 1
}
*/
