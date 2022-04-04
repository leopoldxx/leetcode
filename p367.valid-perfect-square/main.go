package main

func sqrt(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	y := num
	yy := num / y
	for y > yy {
		y = (y + yy) / 2
		yy = num / y
	}
	return y
}

func isPerfectSquare(num int) bool {
	y := sqrt(num)
	return y*y == num
}
