package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 1
	}
	bs := []byte(s)
	counter := [256]int{}

	maxLen := 0
	left, right := 0, 0
	for right < len(s) {
		counter[int(bs[right])]++
		for counter[int(bs[right])] > 1 {
			counter[int(bs[left])]--
			left++
		}
		right++
		currentLen := right - left
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}
