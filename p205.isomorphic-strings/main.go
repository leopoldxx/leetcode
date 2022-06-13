package main

func isIsomorphic(s string, t string) bool {
	fm := map[byte]byte{}
	bm := map[byte]byte{}

	if len(s) != len(t) {
		return false
	}
	n := len(s)
	for i := 0; i < n; i++ {
		sc := s[i]
		tc := t[i]
		_, ok := fm[sc]
		_, ok2 := bm[tc]
		if !ok && !ok2 {
			fm[sc] = tc
			bm[tc] = sc
			continue
		}
		if fm[sc] != tc || bm[tc] != sc {
			return false
		}
	}
	return true
}
