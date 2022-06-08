package main

func sumOfArray(nums []int) (res int) {
	for _, num := range nums {
		res += num
	}
	return
}

func indexOfNearestSum(nums []int, x int) (int, int) {
	sum := 0
	for j := len(nums) - 1; j >= 0; j-- {
		sum += nums[j]
		if sum >= x {
			return j, sum
		}
	}
	return -1, sum
}

func minOperations(nums []int, x int) int {
	sum := sumOfArray(nums)
	if sum < x {
		return -1
	} else if sum == x {
		return len(nums)
	}
	idx, sum := indexOfNearestSum(nums, x)
	if idx < 0 {
		return -1
	}
	var result *int
	if sum == x {
		length := len(nums) - idx
		result = &length
	}
	left := 0
	for idx < len(nums) {
		sum -= nums[idx]
		idx++
		for sum < x {
			sum += nums[left]
			left++
		}
		if sum == x {
			if result == nil {
				length := len(nums) - idx + left
				result = &length
			} else if *result > len(nums)-idx+left {
				*result = len(nums) - idx + left
			}
		}
	}
	if result == nil {
		return -1
	}
	return *result
}
