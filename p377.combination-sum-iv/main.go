package main

import "sort"

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	sort.Ints(nums)
	dp[0] = 1
	for current := 1; current <= target; current++ {
		sum := 0
		for i := 0; i < len(nums); i++ {
			if diff := current - nums[i]; diff >= 0 {
				sum += dp[diff]
				continue
			}
			break
		}
		dp[current] = sum
	}
	return dp[target]

}
