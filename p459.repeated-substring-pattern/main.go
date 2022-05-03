package main

func repeatedSubstringPattern(s string) bool {
	bs := []byte(s)
	nbs := len(bs)
	for n := 1; n < nbs; n++ {
		if (nbs-n)%n != 0 {
			continue
		}
		start := 0
		equal := true
		for i := n; i < nbs; i++ {
			if bs[i] != bs[start] {
				equal = false
				break
			}
			start = (start + 1) % n
		}
		if equal {
			return true
		}
	}
	return false
}
