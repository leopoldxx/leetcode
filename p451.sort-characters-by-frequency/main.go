package main

import (
	"sort"
	"strings"
)

func frequencySort(s string) string {
	bs := []byte(s)
	counter := make([][2]int, 128)

	for _, b := range bs {
		counter[int(b)] = [2]int{int(b), counter[int(b)][1] + 1}
	}
	sort.Slice(counter, func(i, j int) bool {
		return counter[i][1] > counter[j][1]
	})
	sb := strings.Builder{}
	sb.Grow(len(s))
	for i := range counter {
		if counter[i][1] == 0 {
			break
		}
		for j := 0; j < counter[i][1]; j++ {
			sb.WriteByte(byte(counter[i][0]))
		}
	}
	return sb.String()
}
