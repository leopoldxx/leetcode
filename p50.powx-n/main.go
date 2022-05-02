package main

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func myPow(x float64, n int) float64 {
	var result float64 = 1
	absn := abs(n)

	xpow2 := x
	for absn > 0 {
		if absn%2 == 1 {
			result *= xpow2
		}
		xpow2 *= xpow2
		absn /= 2
	}

	if n < 0 {
		return 1 / result
	}
	return result
}
