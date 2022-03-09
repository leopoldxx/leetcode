package main

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if currentSum < 0 {
			currentSum = num
		} else {
			currentSum = currentSum + num
		}
		if maxSum < currentSum {
			maxSum = currentSum
		}
	}
	return maxSum

}
