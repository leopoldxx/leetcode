package main

var letters = [][]byte{
	{},                   // 0
	{},                   // 1
	{'a', 'b', 'c'},      // 2
	{'d', 'e', 'f'},      // 3
	{'g', 'h', 'i'},      // 4
	{'j', 'k', 'l'},      // 5
	{'m', 'n', 'o'},      // 6
	{'p', 'q', 'r', 's'}, // 7
	{'t', 'u', 'v'},      // 8
	{'w', 'x', 'y', 'z'}, // 9
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	byted := []byte(digits)
	pre := [][]byte{}
	for _, l := range letters[byted[0]-'0'] {
		pre = append(pre, []byte{l})
	}

	for i := 1; i < len(byted); i++ {
		res := [][]byte{}
		bs := letters[byted[i]-'0']
		for j := 0; j < len(pre); j++ {
			for n := 0; n < len(bs); n++ {
				res = append(res, append(append([]byte(nil), pre[j]...), bs[n]))
			}
		}
		pre = res
	}

	var result []string
	for i := 0; i < len(pre); i++ {
		result = append(result, string(pre[i]))
	}
	return result

}
