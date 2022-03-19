package main

func isValidCode(c []byte) bool {
	if len(c) == 1 {
		return c[0] >= '1' && c[0] <= '9'
	} else if len(c) == 2 {
		if c[0] != '1' && c[0] != '2' {
			return false
		} else if c[0] == '2' && c[1] > '6' {
			return false
		}
		return true
	}
	return false
}
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func numDecodings(s string) int {
	bs := []byte(s)
	if len(bs) == 1 {
		if isValidCode(bs) {
			return 1
		}
		return 0
	} else if len(bs) == 2 {
		return boolToInt(isValidCode(bs)) + boolToInt(isValidCode(bs[:1]) && isValidCode(bs[1:]))
	}
	dp := make([]int, len(s))
	dp[0] = boolToInt(isValidCode(bs[:1]))
	dp[1] = boolToInt(isValidCode(bs[:2])) + boolToInt(isValidCode(bs[:1]) && isValidCode(bs[1:2]))

	if dp[0]+dp[1] == 0 {
		return 0
	}

	for i := 2; i < len(s); i++ {
		dp[i] = 0
		if isValidCode(bs[i:i+1]) && dp[i-1] > 0 {
			dp[i] += dp[i-1]
		}
		if isValidCode(bs[i-1:i+1]) && dp[i-2] > 0 {
			dp[i] += dp[i-2]
		}
		if dp[i] == 0 && dp[i-1] == 0 {
			return 0
		}
	}
	return dp[len(s)-1]
}
