package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	dp := map[string][]string{}
	for i := range strs {
		bs := []byte(strs[i])
		sort.Slice(bs, func(i, j int) bool {
			if bs[i] < bs[j] {
				return true
			}
			return false
		})
		dp[string(bs)] = append(dp[string(bs)], strs[i])
	}
	result := [][]string{}
	for _, v := range dp {
		result = append(result, v)
	}
	return result
}
