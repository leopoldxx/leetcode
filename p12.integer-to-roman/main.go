package main

func intToRoman(num int) string {
	res := make([]byte, 0, 10)
	for num > 0 {
		switch {
		case num >= 1000:
			c := num / 1000
			for i := 0; i < c; i++ {
				res = append(res, 'M')
			}
			num -= c * 1000

		case num >= 900:
			res = append(res, 'C')
			res = append(res, 'M')
			num -= 900

		case num >= 500:
			res = append(res, 'D')
			num -= 500

			c := num / 100
			for i := 0; i < c; i++ {
				res = append(res, 'C')
			}
			num -= c * 100

		case num >= 400:
			res = append(res, 'C')
			res = append(res, 'D')
			num -= 400

		case num >= 100:
			c := num / 100
			for i := 0; i < c; i++ {
				res = append(res, 'C')
			}
			num -= c * 100

		case num >= 90:
			res = append(res, 'X')
			res = append(res, 'C')
			num -= 90

		case num >= 50:
			res = append(res, 'L')
			num -= 50

			c := num / 10
			for i := 0; i < c; i++ {
				res = append(res, 'X')
			}
			num -= c * 10

		case num >= 40:
			res = append(res, 'X')
			res = append(res, 'L')
			num -= 40

		case num >= 10:
			c := num / 10
			for i := 0; i < c; i++ {
				res = append(res, 'X')
			}
			num -= c * 10

		case num >= 9:
			res = append(res, 'I')
			res = append(res, 'X')
			num -= 90

		case num >= 5:
			res = append(res, 'V')
			num -= 5

			for i := 0; i < num; i++ {
				res = append(res, 'I')
			}
			num -= num

		case num >= 4:
			res = append(res, 'I')
			res = append(res, 'V')
			num -= 4

		case num >= 1:
			for i := 0; i < num; i++ {
				res = append(res, 'I')
			}
			num -= num
		}

	}
	return string(res)
}
