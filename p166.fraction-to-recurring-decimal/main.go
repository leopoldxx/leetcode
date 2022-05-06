package main

import "strconv"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
func itob(n int) []byte {
	return []byte(strconv.Itoa(n))
}
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	result := make([]byte, 0, 10000)

	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
		result = append(result, '-')
	}

	numerator = abs(numerator)
	denominator = abs(denominator)

	result = append(result, itob(numerator/denominator)...)

	numerator %= denominator

	if numerator == 0 {
		return string(result)
	}

	result = append(result, '.')

	poshash := map[int]int{}
	poshash[numerator] = len(result)

	for numerator != 0 {
		numerator *= 10
		result = append(result, itob(numerator/denominator)...)
		numerator %= denominator
		if pos, exists := poshash[numerator]; exists {
			return string(result[:pos]) + "(" + string(result[pos:]) + ")"
		}
		poshash[numerator] = len(result)
	}
	return string(result)
}
