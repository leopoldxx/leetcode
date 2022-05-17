package main

import "sort"

func combine(candidates []int, index int, prefix []int, target int, result *[][]int) {
	if target == 0 {
		*result = append(*result, append([]int(nil), prefix...))
		return
	}

	for i := index; i < len(candidates); i++ {
		if i != index && candidates[i] == candidates[i-1] {
			continue // skip the
		}
		if candidates[i] > target {
			break
		}
		prefix = append(prefix, candidates[i])
		combine(candidates, i+1, prefix, target-candidates[i], result)
		prefix = prefix[:len(prefix)-1]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	result := [][]int{}
	prefix := []int{}
	combine(candidates, 0, prefix, target, &result)
	return result
}
