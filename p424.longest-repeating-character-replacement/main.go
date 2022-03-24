package main

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func characterReplacement(s string, k int) int {
	bs := []byte(s)
	dp := [26]int{}
	left, right, maxUniqCount, maxSubstrLen := 0, 0, 0, 0

	for right = 0; right < len(bs); right++ {
		idx := bs[right] - 'A'
		dp[idx]++
		maxUniqCount = max(maxUniqCount, dp[idx])
		for right-left+1 > maxUniqCount+k {
			dp[bs[left]-'A']--
			left++
		}
		maxSubstrLen = max(maxSubstrLen, right-left+1)
	}
	return maxSubstrLen
}
