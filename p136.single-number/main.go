package main

func singleNumber(nums []int) int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}
	return xor
}
