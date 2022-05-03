package main

const (
	base  = 256
	prime = 101
)

func hashnew(h int, pow int, a, b byte) int {
	return ((h+prime-int(a)*pow%prime)*base + int(b)) % prime
}

func hash(s string) (h int) {
	bs := []byte(s)
	h = int(bs[0]) % prime
	for i := 1; i < len(bs); i++ {
		h = ((h*base)%prime + int(bs[i])) % prime
	}
	return h
}
func hashpow(c int) (pow int) {
	pow = 1
	for i := 0; i < c-1; i++ {
		pow = (pow * base) % prime
	}
	return pow
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	} else if len(needle) > len(haystack) {
		return -1
	}
	pow := hashpow(len(needle))
	hn := hash(needle)
	hh := hash(haystack[:len(needle)])

	hhb := []byte(haystack)

	nh := len(haystack)
	nn := len(needle)

	i := 0
	for i <= nh-nn {
		if hn == hh && needle == haystack[i:i+nn] {
			return i
		}
		i++
		if i > nh-nn {
			return -1
		}
		pre := hhb[i-1]
		newLast := hhb[i+nn-1]
		hh = hashnew(hh, pow, pre, newLast)
	}
	return -1
}
