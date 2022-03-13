package main

import "sort"

const (
	modulo = 10e9 + 7
)

func foundSubTree(arr []int, countMap map[int]int, target int) int {
	var res int
	for i := 0; i < len(arr); i++ {
		if target%arr[i] != 0 {
			continue
		}
		left := countMap[arr[i]]

		if right, exists := countMap[target/arr[i]]; exists {
			res = (res + left*right) % modulo
		}
	}
	return res
}

func numFactoredBinaryTrees(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	sort.Ints(arr)

	countMap := map[int]int{}

	countMap[arr[0]] = 1

	for i := 1; i < len(arr); i++ {
		countMap[arr[i]] = 1 + foundSubTree(arr[:i], countMap, arr[i])
	}
	var res int
	for _, count := range countMap {
		res = (res + count%modulo) % modulo
	}
	return res
}
