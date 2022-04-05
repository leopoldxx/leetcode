package main

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func deleteAndEarn(nums []int) int {
	sums := [10001]int{}
	for _, n := range nums {
		sums[n] += n
	}
	for i := 2; i < 10001; i++ {
		sums[i] = max(sums[i-1], sums[i-2]+sums[i])
	}
	return sums[10000]
}
