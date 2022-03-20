package main

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		currentMaxIndex := i + nums[i]
		if currentMaxIndex > maxIndex {
			maxIndex = currentMaxIndex
		} else if maxIndex == currentMaxIndex && nums[i] == 0 {
			return false
		}

		if maxIndex >= len(nums)-1 {
			return true
		}

	}
	return maxIndex >= len(nums)-1
}
