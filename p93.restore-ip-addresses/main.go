package main

import (
	"fmt"
	"strconv"
)

func isvalid(s string) (v int, ok bool) {
	if len(s) == 0 {
		ok = false
		return
	}
	if s[0] == '0' {
		return 0, s == "0"
	}
	v, _ = strconv.Atoi(s)
	if v < 0 || v > 255 {
		return 0, false
	}
	return v, true
}
func combine(a []int, b int) (res []int) {
	res = append(res, a...)
	res = append(res, b)
	return
}
func dfs(s string, prefix []int, res *[][]int) {
	if len(prefix) == 3 {
		v, ok := isvalid(s)
		if !ok {
			return
		}
		*res = append(*res, combine(prefix, v))
		return
	}

	for i := 1; i <= 3 && i < len(s); i++ {
		si := s[:i]
		v, ok := isvalid(si)
		if !ok {
			break
		}
		prefix = append(prefix, v)
		dfs(s[i:], prefix, res)
		prefix = prefix[:len(prefix)-1]
	}
}

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	res := [][]int{}
	dfs(s, []int{}, &res)

	result := make([]string, 0, len(res))
	for i := 0; i < len(res); i++ {
		result = append(result, fmt.Sprintf("%d.%d.%d.%d", res[i][0], res[i][1], res[i][2], res[i][3]))
	}
	return result
}
