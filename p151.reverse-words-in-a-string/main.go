package main

import "strings"

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	sb := &strings.Builder{}
	for i := len(ss) - 1; i >= 0; i-- {
		if ss[i] == "" {
			continue
		}
		if sb.Len() > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(ss[i])
	}
	return sb.String()
}
