package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}
	if minLen == 0 {
		return ""
	}
	max := 0
	for i := 0; i < minLen; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[0][i] != strs[j][i] {
				return strs[0][:max]
			}
		}
		max++
	}
	return strs[0][:max]
}
