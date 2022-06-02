package main

func encode(s []byte) (res int) {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'A':
			res = res*10 + 1
		case 'C':
			res = res*10 + 2
		case 'G':
			res = res*10 + 3
		case 'T':
			res = res*10 + 4
		}
	}
	return
}
func decode(code int) string {
	bs := make([]byte, 10)
	for i := 9; i >= 0; i-- {
		d := code % 10
		switch d {
		case 1:
			bs[i] = 'A'
		case 2:
			bs[i] = 'C'
		case 3:
			bs[i] = 'G'
		case 4:
			bs[i] = 'T'
		}
		code /= 10
	}
	return string(bs)
}

func encodeLite(code int, next byte) int {
	code %= 1000000000
	code = code * 10

	switch next {
	case 'A':
		code += 1
	case 'C':
		code += 2
	case 'G':
		code += 3
	case 'T':
		code += 4
	}
	return code
}

func findRepeatedDnaSequences(s string) []string {
	bs := []byte(s)
	if len(bs) <= 10 {
		return nil
	}
	hash := map[int]int{}
	v := encode(bs[:10])
	hash[v]++
	for i := 10; i < len(bs); i++ {
		v = encodeLite(v, bs[i])
		hash[v]++
	}

	res := []string{}
	for k, c := range hash {
		if c <= 1 {
			continue
		}
		res = append(res, decode(k))
	}
	return res
}
