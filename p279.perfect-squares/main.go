package main

func genPerfectNums(n int) []int {
	square := 1
	res := []int{}
	for i := 1; ; i++ {
		square = i * i
		if square <= n {
			res = append(res, square)
		} else {
			break
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a

	}
	return b
}

func numSquares(n int) int {
	nums := genPerfectNums(n)
	dp := map[int]int{
		0: 0,
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < len(nums) && nums[j] <= i; j++ {
			if v, ok := dp[i-nums[j]]; ok {
				if _, ok2 := dp[i]; !ok2 {
					dp[i] = v + 1
				} else {
					dp[i] = min(v+1, dp[i])
				}
			}
		}
	}
	return dp[n]
}
