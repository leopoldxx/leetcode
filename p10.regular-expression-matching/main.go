package main

func matchAnything(p []byte) bool {
	if len(p)%2 == 1 {
		return false
	}
	for i := 1; i < len(p); i += 2 {
		if p[i] != '*' {
			return false
		}
	}
	return true
}
func match(s []byte, p []byte, si, pi int, dp map[[2]int]bool) bool {
	if pi == len(p) {
		return si == len(s)
	} else if si == len(s) {
		return matchAnything(p[pi:])
	}
	if v, ok := dp[[2]int{si, pi}]; ok {
		return v
	}

	matchByte := s[si] == p[pi] || p[pi] == '.'
	if pi+1 < len(p) && p[pi+1] == '*' {
		m1 := match(s, p, si, pi+2, dp)
		dp[[2]int{si, pi + 2}] = m1
		m2 := matchByte && match(s, p, si+1, pi, dp)
		dp[[2]int{si, pi}] = m2

		return m1 || m2
	}
	m := matchByte && match(s, p, si+1, pi+1, dp)
	dp[[2]int{si, pi}] = m
	return m
}

func isMatch(s string, p string) bool {
	dp := map[[2]int]bool{}
	return match([]byte(s), []byte(p), 0, 0, dp)
}
