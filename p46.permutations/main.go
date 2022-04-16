package main

func permuteWithAuxTakenArray(nums []int, prefix []int, taken [6]bool) [][]int {
	if len(prefix) == len(nums) {
		return [][]int{prefix}
	}
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if taken[i] {
			continue
		}
		taken[i] = true
		p := permuteWithAuxTakenArray(nums, append(prefix, nums[i]), taken)
		taken[i] = false
		res = append(res, p...)
	}
	return res
}

func permute(nums []int) [][]int {
	taken := [6]bool{}

	return permuteWithAuxTakenArray(nums, []int{}, taken)
}
