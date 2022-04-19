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
			nums:   []int{1, 4, 2},
			result: 3,
		},
		{
			nums:   []int{1, 2, 3, 0, 2},
			result: 3,
		},
		{
			nums:   []int{1},
			result: 0,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, maxProfit(tc.nums))
	}
}
