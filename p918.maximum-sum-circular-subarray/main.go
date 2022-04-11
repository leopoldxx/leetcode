package main

func max2(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func max3(a, b, c int) int {
	if a < 0 {
		return c
	}
	return b
}

func min3(a, b, c int) int {
	if a > 0 {
		return c
	}
	return b
}

func maxSubarraySumCircular(nums []int) int {
	total, maxSum, minSum, currentMaxSum, currentMinSum := nums[0], nums[0], nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		total += nums[i]
		currentMaxSum = max3(currentMaxSum, currentMaxSum+nums[i], nums[i])
		currentMinSum = min3(currentMinSum, currentMinSum+nums[i], nums[i])
		if maxSum < currentMaxSum {
			maxSum = currentMaxSum
		}
		if minSum > currentMinSum {
			minSum = currentMinSum
		}
	}
	if minSum > 0 {
		return total
	}
	if maxSum < 0 {
		return maxSum
	}

	return max2(maxSum, total-minSum)

}
