package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var (
		maxLen = 0
		filter = [256]int{}
		left   = 0
		right  = 0
		r      rune
	)

	for right, r = range s {
		filter[int(r)] = filter[int(r)] + 1
		if filter[int(r)] > 1 {
			for left < right {
				filter[int(s[left])] = filter[int(s[left])] - 1
				left++
				if filter[int(r)] == 1 {
					break
				}
			}
		} else {
			newLen := right - left + 1
			if maxLen < newLen {
				maxLen = newLen
			}
		}
	}
	return maxLen
}
