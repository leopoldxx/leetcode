package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums   []int
		result int
	}{
		{
			nums:   []int{17, 0, 17, 0, 5, -10, -15, 13, -14, -3},
			result: 6,
		},
		{
			nums:   []int{70, -18, 75, -72, -69, -84, 64, -65, 0, -82, 62, 54, -63, -85, 53, -60, -59, 29, 32, 59, -54, -29, -45, 0, -10, 22, 42, -37, -16, 0, -7, -76, -34, 37, -10, 2, -59, -24, 85, 45, -81, 56, 86},
			result: 14,
		},
		{
			nums:   []int{100},
			result: 1,
		},
		{
			nums:   []int{0, 0, 0, 0, 0},
			result: 0,
		},
		{
			nums:   []int{0, 1, -2, -3, -4},
			result: 3,
		},
		{
			nums:   []int{-1, 2},
			result: 1,
		},
		{
			nums:   []int{1, -2, -3, 4},
			result: 4,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, getMaxLen(tc.nums))
	}
}
