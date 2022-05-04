package main

func f(n int, cache [21]int) int {
	if cache[n] > 0 {
		return cache[n]
	}
	sum := 0
	for i := 1; i <= n; i++ {
		if i == 1 || i == n {
			sum += f(n-1, cache)
			continue
		}
		sum += f(i-1, cache) * f(n-i, cache)
	}
	cache[n] = sum
	return sum
}

func numTrees(n int) int {
	cache := [21]int{
		0,
		1,
		2,
		5,
		14,
		42,
		132,
		429,
		1430,
		4862,
		16796,
		58786,
		208012,

		// 0: 0,
		// 1: 1,
		// 2: 2,
		// 3: 5,
		// 4: 14,
		// 5: 42,
		// 6: 132,
		// 7: 429,
		// 8: 1430,
		// 9: 4862,
		// 10: 16796,
		// 11: 58786,
		// 12: 208012,
	}
	return f(n, cache)
}
