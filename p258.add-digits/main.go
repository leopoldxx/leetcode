package main

func addDigits(num int) int {
	for num >= 10 {
		newNum := 0
		for num > 0 {
			newNum += (num % 10)
			num /= 10
		}
		num = newNum
	}
	return num
}
