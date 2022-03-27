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
			nums:   []int{0, 1, 2},
			result: 3,
		},
		{
			nums:   []int{3, 4, -1, 1},
			result: 2,
		},
		{
			nums:   []int{7, 8, 9, 11, 12},
			result: 1,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, firstMissingPositive(tc.nums))
	}
}
