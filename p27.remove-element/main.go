package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 && nums[0] == val {
		return 0
	} else if len(nums) == 1 && nums[0] != val {
		return 1
	}
	inplace := nums
	idx := 0
	for {
		if idx == len(inplace) {
			break
		}
		if inplace[idx] == val {
			inplace[idx] = inplace[len(inplace)-1]
			inplace = inplace[:len(inplace)-1]
			continue
		}
		idx++
	}
	return idx
}
