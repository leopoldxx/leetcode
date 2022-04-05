package main

func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func reverseWords(s string) string {
	bs := []byte(s)
	left := 0
	right := 1
	for right < len(bs) {
		if bs[right] == ' ' {
			reverseString(bs[left:right])
			left = right + 1
		}
		right++
	}
	if left != right {
		reverseString(bs[left:right])
	}
	return string(bs)
}
