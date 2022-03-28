package main

func longestConsecutive(nums []int) int {
	dp := map[int]struct{}{}
	for _, num := range nums {
		dp[num] = struct{}{}
	}
	max := 0
	for num := range dp {
		if _, exists := dp[num-1]; exists {
			continue
		}
		right := num + 1
		for {
			_, exists := dp[right]
			if !exists {
				break
			}
			right = right + 1
		}
		if max < right-num {
			max = right - num
		}
	}
	return max
}
