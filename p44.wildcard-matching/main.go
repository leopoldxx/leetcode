package main

func matchAnySeq(p []byte) bool {
	for i := 0; i < len(p); i++ {
		if p[i] != '*' {
			return false
		}
	}
	return true
}

func match(s, p []byte, si, pi int, dp map[[2]int]bool) bool {
	if pi == len(p) {
		return si == len(s)
	} else if si == len(s) {
		return matchAnySeq(p[pi:])
	}

	if v, ok := dp[[2]int{si, pi}]; ok {
		return v
	}

	// match byte or match ?
	if p[pi] != '*' {
		m1 := (s[si] == p[pi] || p[pi] == '?') && match(s, p, si+1, pi+1, dp)
		dp[[2]int{si, pi}] = m1
		return m1
	}

	// match *
	m2 := match(s, p, si+1, pi, dp)
	dp[[2]int{si + 1, pi}] = m2

	m3 := match(s, p, si+1, pi+1, dp)
	dp[[2]int{si + 1, pi + 1}] = m3

	m4 := match(s, p, si, pi+1, dp)
	dp[[2]int{si, pi + 1}] = m4

	dp[[2]int{si, pi}] = m2 || m3 || m4
	return dp[[2]int{si, pi}]
}

func isMatch(s string, p string) bool {
	dp := map[[2]int]bool{}
	return match([]byte(s), []byte(p), 0, 0, dp)
}
