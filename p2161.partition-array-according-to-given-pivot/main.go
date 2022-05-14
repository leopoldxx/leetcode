package main

func switchArrayElement(nums []int, pivot int) int {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > pivot && nums[i+1] <= pivot {
			nums[i], nums[i+1] = nums[i+1], nums[i]
			count++
		} else if nums[i] == pivot && nums[i+1] < pivot {
			nums[i], nums[i+1] = nums[i+1], nums[i]
			count++
		}
	}
	return count
}
func pivotArrayLE1000(nums []int, pivot int) []int {
	switchCount := switchArrayElement(nums, pivot)
	for switchCount > 0 {
		switchCount = switchArrayElement(nums, pivot)
	}
	return nums
}

func pivotArrayWithTemp(nums []int, pivot int) []int {
	lower, equal := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == pivot {
			equal++
		} else if nums[i] < pivot {
			lower++
		}
	}
	lowerIndex := 0
	equalIndex := lower
	higherIndex := lower + equal
	tmp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] == pivot {
			tmp[equalIndex] = nums[i]
			equalIndex++
		} else if nums[i] < pivot {
			tmp[lowerIndex] = nums[i]
			lowerIndex++
		} else {
			tmp[higherIndex] = nums[i]
			higherIndex++
		}
	}
	return tmp
}

func pivotArray(nums []int, pivot int) []int {
	if len(nums) <= 1000 {
		return pivotArrayLE1000(nums, pivot)
	}
	return pivotArrayWithTemp(nums, pivot)
}
