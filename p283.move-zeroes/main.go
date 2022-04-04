package main

func moveZeroes(nums []int) {
	if len(nums) == 1 {
		return
	}
	zeroPointer := 0
	nonZeroPointer := 1
	for nonZeroPointer < len(nums) && zeroPointer < len(nums) {
		for zeroPointer < len(nums) && nums[zeroPointer] != 0 {
			zeroPointer++
		}
		if zeroPointer == len(nums) {
			break
		}
		if nonZeroPointer < zeroPointer {
			nonZeroPointer = zeroPointer + 1
		}
		for nonZeroPointer < len(nums) && nums[nonZeroPointer] == 0 {
			nonZeroPointer++
		}
		if nonZeroPointer == len(nums) {
			break
		}
		nums[zeroPointer], nums[nonZeroPointer] = nums[nonZeroPointer], nums[zeroPointer]
		zeroPointer++
		nonZeroPointer++
	}

}
