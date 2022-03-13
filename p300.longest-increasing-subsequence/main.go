package main

func lengthOfLIS(nums []int) int {
	dp := [2500]int{}
	finalMax := 0
	for i, num := range nums {
		dp[i] = 1

		preMax := 0

		for j := i - 1; j >= preMax; j-- {
			if nums[j] < num && preMax < dp[j] {
				preMax = dp[j]
			}
		}
		dp[i] = dp[i] + preMax
		if finalMax < dp[i] {
			finalMax = dp[i]
		}
	}
	return finalMax
}
