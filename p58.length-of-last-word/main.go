package main

func lengthOfLastWord(s string) int {
	bs := []byte(s)
	length := 0
	counting := false
	for j := len(bs) - 1; j >= 0; j-- {
		if bs[j] == ' ' && !counting {
			continue
		} else if bs[j] == ' ' && counting {
			return length
		}
		counting = true
		length++
	}
	return length
}
