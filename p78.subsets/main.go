package main

func dfs(nums []int, index int, prefix []int, result *[][]int) {
	prefix = append([]int{nums[index]}, prefix...)
	*result = append(*result, prefix)

	for j := index + 1; j < len(nums); j++ {
		dfs(nums, j, prefix, result)
	}
}

func subsets(nums []int) [][]int {
	result := [][]int{{}}
	for i := 0; i < len(nums); i++ {
		prefix := []int{}
		dfs(nums, i, prefix, &result)
	}
	return result
}
