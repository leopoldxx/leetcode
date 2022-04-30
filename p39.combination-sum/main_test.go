package main

import (
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func intSlicesToStringSlice(numss [][]int) (res []string) {
	for _, nums := range numss {
		sort.Ints(nums)
		sb := &strings.Builder{}
		for _, num := range nums {
			sb.WriteString(strconv.Itoa(num))
			sb.WriteString(",")
		}
		res = append(res, sb.String())
	}
	sort.Strings(res)
	return res
}

func TestSolution(t *testing.T) {
	cases := []struct {
		candidates []int
		target     int
		result     [][]int
	}{
		{
			candidates: []int{2, 7, 6, 3, 5, 1},
			target:     9,
			result: [][]int{
				{1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 2},
				{1, 1, 1, 1, 1, 1, 3},
				{1, 1, 1, 1, 1, 2, 2},
				{1, 1, 1, 1, 2, 3},
				{1, 1, 1, 1, 5},
				{1, 1, 1, 2, 2, 2},
				{1, 1, 1, 3, 3},
				{1, 1, 1, 6},
				{1, 1, 2, 2, 3},
				{1, 1, 2, 5},
				{1, 1, 7},
				{1, 2, 2, 2, 2},
				{1, 2, 3, 3},
				{1, 2, 6},
				{1, 3, 5},
				{2, 2, 2, 3},
				{2, 2, 5},
				{2, 7},
				{3, 3, 3},
				{3, 6},
			},
		},
	}
	for _, tc := range cases {
		assert.ElementsMatch(t, intSlicesToStringSlice(tc.result), intSlicesToStringSlice(combinationSum(tc.candidates, tc.target)))
		//assert.(t, tc.result, combinationSum(tc.candidates, tc.target))
	}
}
