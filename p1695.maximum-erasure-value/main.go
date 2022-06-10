package main

func maximumUniqueSubarray(nums []int) int {
	left := 0
	right := 0

	dp := map[int]int{}
	sum := 0
	maxSum := 0
	for right < len(nums) {
		n := nums[right]
		dp[n]++
		sum += n
		for dp[n] > 1 && left < len(nums) {
			m := nums[left]
			sum -= m
			dp[m]--
			left++
		}
		if sum > maxSum {
			maxSum = sum
		}
		right++
	}
	return maxSum
}
