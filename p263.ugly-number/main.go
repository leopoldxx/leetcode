package main

func isUgly(n int) bool {
	if n < 0 {
		return false
	}
	for {
		if n == 1 || n == -1 {
			return true
		}
		p := n
		if n%2 == 0 {
			n = n / 2
		}
		if n%3 == 0 {
			n = n / 3
		}
		if n%5 == 0 {
			n = n / 5
		}
		if p == n {
			return false
		}
	}
}
