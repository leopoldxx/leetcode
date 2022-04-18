package main

func isHappy(n int) bool {
	cache := map[int]bool{}

	f := func(n int) int {
		var result int
		for n > 0 {
			mod := n % 10
			result += mod * mod
			n = n / 10
		}
		return result
	}

	for {
		n = f(n)
		if n == 1 {
			return true
		}
		if cache[n] {
			return false
		}
		cache[n] = true
	}
}
