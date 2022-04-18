package main

func titleToNumber(columnTitle string) int {
	count := func(c byte) int {
		return int(c - 'A' + 1)
	}
	bs := []byte(columnTitle)
	var result int
	for i := 0; i < len(bs); i++ {
		n := count(bs[0])
		if result == 0 {
			result = n
		} else {
			result = result*26 + n
		}
	}
	return result

}
