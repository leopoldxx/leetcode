package main

func max(l, r int) int {
	if l < r {
		return r
	}
	return l
}

func maxamountInto(l, r amount, current int, result *amount) {
	result.amountWith1 = r.amountWith1
	result.includeCurrentWith1 = false
	if l.amountWith1+current > r.amountWith1 {
		result.amountWith1 = l.amountWith1 + current
		result.includeCurrentWith1 = true
	}
	result.amountWithout1 = max(l.amountWithout1+current, r.amountWithout1)
}

type amount struct {
	amountWith1         int
	includeCurrentWith1 bool
	amountWithout1      int
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	} else if len(nums) == 3 {
		return max(max(nums[0], nums[1]), nums[2])
	}
	dp := make([]amount, len(nums))
	dp[0] = amount{amountWith1: nums[0], includeCurrentWith1: true, amountWithout1: 0}
	dp[1] = amount{amountWith1: nums[0], includeCurrentWith1: false, amountWithout1: nums[1]}
	dp[2] = amount{amountWith1: nums[0] + nums[2], includeCurrentWith1: true, amountWithout1: max(nums[1], nums[2])}

	for i := 3; i < len(nums); i++ {
		maxamountInto(dp[i-2], dp[i-1], nums[i], &dp[i])
	}

	last := len(nums) - 1
	if dp[last].includeCurrentWith1 {
		return max(dp[last-1].amountWith1, dp[last].amountWithout1)
	}
	return max(dp[last].amountWith1, dp[last].amountWithout1)
}
