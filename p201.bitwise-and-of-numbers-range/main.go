package main

func rangeBitwiseAnd(left int, right int) int {
	for left < right {
		right = right & (right - 1)
	}
	return right
}
