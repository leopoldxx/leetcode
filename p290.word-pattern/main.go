package main

import "strings"

func wordPattern(pattern string, s string) bool {
	items := strings.Split(s, " ")
	dp := map[byte]string{}
	dpM := map[string]byte{}
	if len(pattern) != len(items) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		r := pattern[i]
		if _, ok := dp[r]; !ok {
			if _, ok := dpM[items[i]]; ok {
				return false
			}
			dp[r] = items[i]
			dpM[items[i]] = r
		} else if dp[r] != items[i] {
			return false
		}
	}
	return true
}
