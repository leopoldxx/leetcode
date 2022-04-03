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
			nums:   []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			result: 6,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, minCostClimbingStairs(tc.nums))
	}
}
