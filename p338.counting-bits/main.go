package main

func countBits(n int) []int {
	result := make([]int, n+1)
	result[0] = 0
	for i := 1; i <= n; i++ {
		result[i] = result[i>>1] + (i & 1)
	}
	return result
}

/*
func countBitsOfNum(num int) int {
	count := 0
	for num > 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
func countBits(n int) []int {
	result := make([]int, 0, n+1)
	for i := 0; i <= n; i++ {
		result = append(result, countBitsOfNum(i))
	}
	return result
}
*/
