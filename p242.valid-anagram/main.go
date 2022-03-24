package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cache := map[byte]int{}
	bs := []byte(s)
	bt := []byte(t)
	for i := 0; i < len(bs); i++ {
		cache[bs[i]]++
		cache[bt[i]]--
	}
	for _, v := range cache {
		if v != 0 {
			return false
		}

	}
	return true
}
