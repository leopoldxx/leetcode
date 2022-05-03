package main

func findLeftMin(nums []int) int {
	min := nums[0]
	for i := 0; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}
	}
	return min
}
func findRightMax(nums []int) int {
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}
func findUnsortedSubarray(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	leftmax := nums[0]
	leftmin := 0
	nn := len(nums)
	for i := 1; i < nn; i++ {
		if leftmax <= nums[i] {
			leftmax = nums[i]
			if i == nn-1 {
				return 0
			}
		} else {
			leftmin = findLeftMin(nums[i:])
			break
		}
	}
	leftIndex := 0
	for i := 0; i < nn; i++ {
		if leftmin >= nums[i] {
			continue
		}
		leftIndex = i
		break
	}

	rightMin := nums[nn-1]
	rightMax := 0
	for i := nn - 2; i >= 0; i-- {
		if rightMin >= nums[i] {
			rightMin = nums[i]
		} else {
			rightMax = findRightMax(nums[:i+1])
			break
		}
	}
	rightIndex := 0
	for i := nn - 1; i >= 0; i-- {
		if rightMax <= nums[i] {
			continue
		}
		rightIndex = i
		break
	}

	return rightIndex - leftIndex + 1

}
