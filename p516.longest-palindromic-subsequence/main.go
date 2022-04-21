package main

func max(l, r int) int {
	if l < r {
		return r
	}
	return l
}

func longestCommonSubsequence(row string, col string) int {
	var (
		preCommonCount     [1001]int
		currentCommonCount [1001]int
	)
	for _, c1 := range row {
		for j, c2 := range col {
			if c1 == c2 {
				currentCommonCount[j+1] = preCommonCount[j] + 1
			} else {
				currentCommonCount[j+1] = max(currentCommonCount[j], preCommonCount[j+1])
			}
		}
		preCommonCount, currentCommonCount = currentCommonCount, preCommonCount
	}
	return preCommonCount[len(col)]
}

func reverse(s string) string {
	bs := []byte(s)
	length := len(bs)
	rs := make([]byte, length)

	for i := 0; i < len(bs); i++ {
		rs[i] = bs[length-i-1]
	}
	return string(rs)
}

func longestPalindromeSubseq(s string) int {
	return longestCommonSubsequence(s, reverse(s))
}
