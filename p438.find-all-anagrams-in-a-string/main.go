package main

func isAnagram(bs []byte, bt []byte) bool {
	if len(bs) != len(bt) {
		return false
	}
	cache := [26]int{}
	for i := 0; i < len(bs); i++ {
		cache[int(bs[i]-'a')]++
		cache[int(bt[i]-'a')]--
	}
	for _, v := range cache {
		if v != 0 {
			return false
		}
	}
	return true
}

func computeHash(s []byte) int {
	hash := 0
	for _, b := range s {
		hash += int(b - 'a')
	}
	return hash
}

func computeHashIncr(v int, pre byte, next byte) int {
	return v - int(pre-'a') + int(next-'a')
}

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	bp := []byte(p)
	bs := []byte(s)
	np := len(bp)
	ns := len(bs)

	ph := computeHash(bp)
	ps := computeHash(bs[:np])

	res := []int{}
	if ph == ps && isAnagram(bp, bs[:np]) {
		res = append(res, 0)
	}

	for i := 1; i <= ns-np; i++ {
		ps = computeHashIncr(ps, bs[i-1], bs[i+np-1])
		if ph == ps && isAnagram(bp, bs[i:i+np]) {
			res = append(res, i)
			continue
		}
	}
	return res
}
