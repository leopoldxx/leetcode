package main

func count(str string) (zerosCount int, onesCount int) {
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			zerosCount++
		} else {
			onesCount++
		}
	}
	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for k := 0; k < len(strs); k++ {
		zc, oc := count(strs[k])
		for i := m; i >= zc; i-- {
			for j := n; j >= oc; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zc][j-oc]+1)
			}
		}
	}
	return dp[m][n]
}
