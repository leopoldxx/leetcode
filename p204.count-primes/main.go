package main

func countPrimes(n int) int {
	nums := make([]bool, n)
	count := 0

	for i := 2; i < n; i++ {
		if nums[i] {
			continue
		}
		count++
		for j := 2; j*i < n; j++ {
			nums[i*j] = true
		}
	}

	return count
}
