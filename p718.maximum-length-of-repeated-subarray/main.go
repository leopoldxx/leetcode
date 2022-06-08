package main

func findLength(nums1 []int, nums2 []int) int {
	n1 := len(nums1)
	n2 := len(nums2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}
	result := 0
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if result < dp[i][j] {
				result = dp[i][j]
			}
		}
	}
	return result

}
