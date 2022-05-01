package main

func compact(s string) string {
	res := make([]byte, 0, len(s))
	bs := []byte(s)
	for i := 0; i < len(bs); i++ {
		c := bs[i]
		if c != '#' {
			res = append(res, c)
		} else {
			if len(res) != 0 {
				res = res[:len(res)-1]
			}
		}
	}
	return string(res)
}

func backspaceCompare(s string, t string) bool {
	return compact(s) == compact(t)
}
