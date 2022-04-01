package main

func plusOne(digits []int) []int {
	carry := 1
	for j := len(digits) - 1; j >= 0; j-- {
		digits[j] += carry
		if digits[j] < 10 {
			return digits
		}
		carry = 1
		digits[j] = 0
	}
	return append([]int{1}, digits...)
}
