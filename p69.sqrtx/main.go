package main

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	y := x
	yy := x / y
	for y > yy {
		y = (y + yy) / 2
		yy = x / y
	}
	return y
}
