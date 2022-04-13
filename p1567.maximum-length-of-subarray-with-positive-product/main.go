package main

type value struct {
	left  int
	right int
	val   int
}

func maxLen(a, b value) value {
	if a.left < b.left {
		return a
	}
	return b
}

func determinMax(max, min, current value) value {
	if current.val == 0 {
		return current
	} else if max.val == 0 {
		return current
	}

	// max is counting
	if max.val > 0 {
		if min.val > 0 {
			max = maxLen(max, min)
		}
		max.val = 1 // the val value is useless, we only care about the sign
		return max
	} else {
		if min.val < 0 {
			return current
		}
		// min is counting
		// swap min and max
		min.val = 1 // the val value is useless, we only care about the sign
		return min
	}
}

func determinMin(max, min, current value) value {
	if current.val == 0 {
		return current
	} else if min.val == 0 {
		return current
	}

	// min is counting
	if min.val < 0 {
		if max.val < 0 {
			min = maxLen(max, min)
		}
		min.val = -1 // the val value is useless, we only care about the sign
		return min
	} else {
		if max.val > 0 {
			return current
		}
		// max is counting
		// swap min and max
		max.val = -1 // the val value is useless, we only care about the sign
		return max
	}
}

func getMaxLen(nums []int) int {
	maxSum := value{left: 0, right: 0, val: nums[0]}
	minSum := value{left: 0, right: 0, val: nums[0]}

	var maxNext, minNext, current value

	var result *value
	if maxSum.val > 0 {
		result = &value{val: maxSum.val, left: maxSum.left, right: maxSum.right}
	}

	for i := 1; i < len(nums); i++ {
		current = value{val: nums[i], left: i, right: i}
		maxNext = value{val: maxSum.val * nums[i], left: maxSum.left, right: i}
		minNext = value{val: minSum.val * nums[i], left: minSum.left, right: i}

		maxSum = determinMax(maxNext, minNext, current)
		minSum = determinMin(maxNext, minNext, current)

		if maxSum.val > 0 {
			if result == nil || result.right-result.left < maxSum.right-maxSum.left {
				result = &value{val: maxSum.val, left: maxSum.left, right: maxSum.right}
			}
		}
	}
	if result == nil {
		return 0
	}
	return result.right - result.left + 1
}
