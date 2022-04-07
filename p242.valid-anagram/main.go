package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cache := [26]int{}
	bs := []byte(s)
	bt := []byte(t)
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
