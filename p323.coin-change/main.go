package main

import (
	"sort"
)

func min(l, r int) int {
	if l < r {
		return l
	}
	return r
}

func minCount(needs map[int]int) int {
	if len(needs) <= 0 {
		return -1
	}
	m := int(^uint(0) >> 1)
	for _, v := range needs {
		m = min(m, v)
	}
	return m
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	sort.Ints(coins)
	dp := [10001]map[int]int{}
	for i := 1; i <= amount; i++ {
		dp[i] = map[int]int{}
		for _, num := range coins {
			needs := i - num
			if needs < 0 {
				break
			} else if needs == 0 {
				dp[i][num] = 1
				continue
			}
			if len(dp[needs]) > 0 {
				dp[i][num] = minCount(dp[needs]) + 1
			}
		}
	}
	return minCount(dp[amount])
}
