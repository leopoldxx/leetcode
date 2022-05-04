package main

func isPalindrome(s []byte) bool {
	i := 0
	j := len(s) - 1
	for i <= j {
		if i == j {
			return true
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

var (
	cache = map[string][][]string{}
)

func partition(s string) [][]string {
	if len(s) == 1 {
		return [][]string{{s}}
	}
	if res, exists := cache[s]; exists {
		return res
	}
	res := [][]string{}

	for l := 1; l < len(s); l++ {
		if !isPalindrome([]byte(s[0:l])) {
			continue
		}
		next := s[l:]
		var (
			subres [][]string
			exists bool
		)
		if subres, exists = cache[next]; !exists {
			subres = partition(next)
		}
		for i := 0; i < len(subres); i++ {
			res = append(res, append([]string{s[0:l]}, subres[i]...))
		}
	}
	if isPalindrome([]byte(s)) {
		res = append(res, []string{s})
	}
	cache[s] = res
	return res
}
