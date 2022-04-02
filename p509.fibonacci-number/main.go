package main

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	n_2 := 0
	n_1 := 1
	s := n_2
	for i := 2; i <= n; i++ {
		s = n_1 + n_2
		n_2 = n_1
		n_1 = s
	}
	return s
}
