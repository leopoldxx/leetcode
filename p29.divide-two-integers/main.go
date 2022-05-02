package main

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

var (
	maxInt32 = 1<<31 - 1
	minInt32 = -1 << 31
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	} else if divisor == 0 {
		return minInt32
	}
	if dividend == int(minInt32) && divisor == -1 {
		return maxInt32
	}

	isNegtive := (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0)
	dividend = abs(dividend)
	divisor = abs(divisor)
	result := 0
	for dividend >= divisor {
		shift := 0
		for dividend >= (divisor << shift) {
			shift++
		}
		dividend -= divisor << (shift - 1)
		result += 1 << (shift - 1)
	}
	if isNegtive {
		return -result
	}
	return result
}
