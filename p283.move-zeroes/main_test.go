package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums   []int
		result []int
	}{
		{
			nums:   []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0},
			result: []int{4, 2, 4, 3, 5, 1, 0, 0, 0, 0},
		},
		{
			nums:   []int{0, 1, 0, 3, 12},
			result: []int{1, 3, 12, 0, 0},
		},
	}
	for _, tc := range cases {
		moveZeroes(tc.nums)
		assert.Equal(t, tc.result, tc.nums)
	}
}
