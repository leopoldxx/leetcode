package main

func distance(a, b float64) float64 {
	d := a - b
	if d < 0 {
		return -d
	}
	return d
}

func sqrt(x float64) float64 {
	if x == 0 || x == 1 {
		return x
	}
	y := x
	yy := x / y
	for distance(y, yy) > 0.0001 {
		y = (y + yy) / float64(2)
		yy = x / y
	}
	return y
}

func arrangeCoins(n int) int {
	return int((sqrt(float64(8*n+1)) - 1) / 2)
}
