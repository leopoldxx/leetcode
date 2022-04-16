package main

func expand(c byte) (byte, byte) {
	if c >= 'a' {
		return c, c - ('a' - 'A')
	}
	return c, c + ('a' - 'A')
}

func permuteWithAuxTakenArray(bs []byte, prefix []byte) []string {
	if len(bs) == 0 {
		return []string{string(prefix)}
	}
	res := []string{}
	c := bs[0]
	if c >= '0' && c <= '9' {
		res = append(res, permuteWithAuxTakenArray(bs[1:], append(prefix, c))...)
	} else {
		c1, c2 := expand(c)
		res = append(res, permuteWithAuxTakenArray(bs[1:], append(prefix, c1))...)
		res = append(res, permuteWithAuxTakenArray(bs[1:], append(prefix, c2))...)
	}
	return res
}

func letterCasePermutation(s string) []string {
	return permuteWithAuxTakenArray([]byte(s), []byte{})
}
