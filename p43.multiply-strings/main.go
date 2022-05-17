package main

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n := len(num1) + len(num2)
	cache := make([]int, n)
	for ii := len(num1) - 1; ii >= 0; ii-- {
		for jj := len(num2) - 1; jj >= 0; jj-- {
			cache[ii+jj+1] += int(num1[ii]-'0') * int(num2[jj]-'0')
		}
	}
	for ii := n - 1; ii > 0; ii-- {
		cache[ii-1] += cache[ii] / 10
		cache[ii] = cache[ii] % 10
	}
	res := []byte{}
	for i := 0; i < n; i++ {
		if cache[i] == 0 && len(res) == 0 {
			continue
		}
		res = append(res, byte(int('0')+cache[i]))
	}
	return string(res)
}
