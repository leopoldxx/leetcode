package main

func addbyte(a, b byte) byte {
	return a - '0' + b
}
func carry(a byte) (byte, int) {
	r := a - '0'
	return byte(r%10) + '0', int(r / 10)
}

func addStrings(num1 string, num2 string) string {
	if num1 == "0" {
		return num2
	} else if num2 == "0" {
		return num1
	}
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	bnum1 := []byte(num1)
	diff := len(num1) - len(num2)

	for ii := len(num2) - 1; ii >= 0; ii-- {
		bnum1[ii+diff] = addbyte(bnum1[ii+diff], num2[ii])
	}

	var c int
	for ii := len(num1) - 1; ii >= 0; ii-- {
		bnum1[ii], c = carry(bnum1[ii] + byte(c))
	}

	if c == 1 {
		return "1" + string(bnum1)
	}
	return string(bnum1)
}
