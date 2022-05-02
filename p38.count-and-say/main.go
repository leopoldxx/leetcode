package main

import (
	"strconv"
	"strings"
)

func compact(v string) string {
	sb := strings.Builder{}
	bs := []byte(v)

	count := 1
	var b byte = bs[0]

	for i := 1; i < len(v); i++ {
		if b == bs[i] {
			count++
		} else {
			sb.WriteString(strconv.Itoa(count))
			sb.WriteByte(b)
			count = 1
			b = bs[i]
		}
	}
	sb.WriteString(strconv.Itoa(count))
	sb.WriteByte(b)

	return sb.String()
}

func countAndSay(n int) string {
	res := "1"
	for i := 1; i < n; i++ {
		res = compact(res)
	}
	return res
}
