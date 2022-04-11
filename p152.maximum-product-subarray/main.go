package main

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}

func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	} else {
		return c
	}
}

func maxProduct(nums []int) int {
	result, maxProd, minProd := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		maxNext := maxProd * nums[i]
		minNext := minProd * nums[i]
		maxProd = max3(maxNext, nums[i], minNext)
		minProd = min3(maxNext, nums[i], minNext)

		if result < maxProd {
			result = maxProd
		}

	}
	return result

}
