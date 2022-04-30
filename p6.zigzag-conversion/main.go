package main

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]byte, 0, len(s)/numRows)
	}
	bs := []byte(s)
	index := 0
	step := 1

	for i := 0; i < len(bs); i++ {
		res[index] = append(res[index], bs[i])
		if index == 0 {
			step = 1
		} else if index == numRows-1 {
			step = -1
		}
		index += step
	}

	sb := strings.Builder{}
	sb.Grow(len(s))
	for i := 0; i < numRows; i++ {
		sb.WriteString(string(res[i]))
	}
	return sb.String()
}
