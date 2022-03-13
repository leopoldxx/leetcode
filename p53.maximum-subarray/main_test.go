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
			nums:   []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			result: 6,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, maxSubArray(tc.nums))
	}
}
