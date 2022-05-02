package main

import (
	"strconv"
)

func encode(b byte, count int) []byte {
	if count == 1 {
		return []byte{b}
	}
	res := make([]byte, 0, 5)
	res = append(res, b)
	res = append(res, []byte(strconv.Itoa(count))...)
	return res
}

func compress(chars []byte) int {
	resultIndex := 0

	count := 1
	b := chars[0]

	for i := 1; i < len(chars); i++ {
		if b == chars[i] {
			count++
		} else {
			res := encode(b, count)
			for j := 0; j < len(res); j++ {
				chars[resultIndex] = res[j]
				resultIndex++
			}

			b = chars[i]
			count = 1
		}
	}
	res := encode(b, count)
	for j := 0; j < len(res); j++ {
		chars[resultIndex] = res[j]
		resultIndex++
	}
	return resultIndex
}
