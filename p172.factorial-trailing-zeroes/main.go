package main

func count5(n int) (count int) {
	for n > 0 && n%5 == 0 {
		count++
		n /= 5
	}
	return
}
func countZero(n int) (count int) {
	for n > 0 && n%10 == 0 {
		count++
		n /= 10
	}
	for n > 0 && n%5 == 0 {
		count++
		n /= 5
	}
	return
}
func trailingZeroes(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if i%10 == 0 {
			count += countZero(i)
		} else if i%5 == 0 {
			count += count5(i)
		}
	}
	return count
}
