package main

func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	}
	n_3 := 0
	n_2 := 1
	n_1 := 1
	s := n_1
	for i := 3; i <= n; i++ {
		s = n_1 + n_2 + n_3
		n_3, n_2, n_1 = n_2, n_1, s
	}
	return s
}
