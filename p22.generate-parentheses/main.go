package main

func generateParenthesisWithArgs(left int, right int) [][]byte {
	if left == 0 {
		combination := make([]byte, right)
		for i := range combination {
			combination[i] = ')'
		}
		return [][]byte{combination}
	}

	var candidates [][]byte
	if left <= right {
		combinations := generateParenthesisWithArgs(left-1, right)
		for i := range combinations {
			candidates = append(candidates, append([]byte{'('}, combinations[i]...))
		}
	}
	if left < right {
		combinations := generateParenthesisWithArgs(left, right-1)
		for i := range combinations {
			candidates = append(candidates, append([]byte{')'}, combinations[i]...))
		}
	}
	return candidates
}

func generateParenthesis(n int) []string {
	bss := generateParenthesisWithArgs(n, n)
	if n == 1 {
		return []string{"()"}
	}
	res := []string{}
	for i := range bss {
		res = append(res, string(bss[i]))
	}
	return res

}
