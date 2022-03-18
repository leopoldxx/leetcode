package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	y := 0
	xx := x
	for x > 0 {
		y = y*10 + x%10
		x /= 10
	}

	return xx == y
}
