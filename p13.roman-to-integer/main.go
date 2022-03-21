package main

func romanToInt(s string) int {
	bs := []byte(s)

	result := 0

	for i := 0; i < len(s); i++ {
		b := bs[i]

		switch b {
		case 'I':
			if i < len(s)-1 && bs[i+1] == 'V' {
				result += 4
				i++
			} else if i < len(s)-1 && bs[i+1] == 'X' {
				result += 9
				i++
			} else {
				result += 1
			}
		case 'V':
			result += 5

		case 'X':
			if i < len(s)-1 && bs[i+1] == 'L' {
				result += 40
				i++
			} else if i < len(s)-1 && bs[i+1] == 'C' {
				result += 90
				i++
			} else {
				result += 10
			}

		case 'L':
			result += 50

		case 'C':
			if i < len(s)-1 && bs[i+1] == 'D' {
				result += 400
				i++
			} else if i < len(s)-1 && bs[i+1] == 'M' {
				result += 900
				i++
			} else {
				result += 100
			}

		case 'D':
			result += 500

		case 'M':
			result += 1000
		}
	}
	return result
}
