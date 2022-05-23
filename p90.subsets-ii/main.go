package main

import "sort"

func dfs(nums []int, index int, prefix []int, result *[][]int) {
	*result = append(*result, prefix)

	for j := index; j < len(nums); j++ {
		if j != index && nums[j] == nums[j-1] {
			continue
		}

		nprefix := append([]int{nums[j]}, prefix...)
		dfs(nums, j+1, nprefix, result)
	}
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	result := make([][]int, 0, n)
	dfs(nums, 0, []int{}, &result)
	return result
}
