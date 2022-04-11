package main

func max(a, b int) int {
	if a < 0 {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > 0 {
		return b
	}
	return a
}

func maxSubarraySumCircular(nums []int) int {
	total, maxSum, minSum, currentMaxSum, currentMinSum := nums[0], nums[0], nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		total += nums[i]
		currentMaxSum = max(currentMaxSum+nums[i], nums[i])
		currentMinSum = min(currentMinSum+nums[i], nums[i])
		if maxSum < currentMaxSum {
			maxSum = currentMaxSum
		}
		if minSum > currentMinSum {
			minSum = currentMinSum
		}
	}

	return max(maxSum, total+minSum)

}
