package main

import "sort"

func permute(nums []int, prefix []int, visited []bool, result *[][]int) {
	if len(prefix) == len(nums) {
		*result = append(*result, prefix)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
			continue
		}
		visited[i] = true
		permute(nums, append([]int{nums[i]}, prefix...), visited, result)
		visited[i] = false
	}

}

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	visited := make([]bool, len(nums))
	sort.Ints(nums)
	result := make([][]int, 0, 100)
	permute(nums, []int{}, visited, &result)
	return result
}
