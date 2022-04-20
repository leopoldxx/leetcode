package main

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	si := 0
	st := 0
	for si < len(s) && st < len(t) {
		if s[si] == t[st] {
			if si == len(s)-1 {
				return true
			}
			si++
			st++
			continue
		}
		st++
	}
	return false
}
