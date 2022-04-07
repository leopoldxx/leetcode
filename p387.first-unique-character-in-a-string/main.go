package main

func firstUniqChar(s string) int {
	letterCount := [26]int{}
	bs := ([]byte)(s)
	for _, b := range bs {
		letterCount[int(b-'a')]++
	}
	for i, b := range bs {
		if letterCount[int(b-'a')] == 1 {
			return i
		}
	}
	return -1
}
