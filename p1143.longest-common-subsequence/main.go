package main

func max(l, r int16) int16 {
	if l < r {
		return r
	}
	return l
}
func longestCommonSubsequence(text1 string, text2 string) int {
	var result [2][1001]int16
	pre, current := 0, 1
	for _, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				result[current][j+1] = result[pre][j] + 1
			} else {
				result[current][j+1] = max(result[current][j], result[pre][j+1])
			}
		}
		current, pre = pre, current
	}
	return int(result[pre][len(text2)])
}
