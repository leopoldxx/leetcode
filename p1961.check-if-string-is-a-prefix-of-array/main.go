package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func isPrefixString(s string, words []string) bool {

	for i := 0; i < len(words); i++ {
		ns := len(s)
		nw := len(words[i])
		if nw > ns {
			return false
		}
		if s[:nw] != words[i] {
			return false
		}
		s = s[nw:]
		if len(s) == 0 {
			return true
		}
	}
	return false
}
