package main

import "strings"

func countPrefixes(words []string, s string) int {
	count := 0
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(s, words[i]) {
			count++
		}
	}
	return count
}
