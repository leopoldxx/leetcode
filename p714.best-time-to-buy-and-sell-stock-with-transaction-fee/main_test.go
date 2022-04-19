package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums   []int
		fee    int
		result int
	}{
		{
			nums:   []int{1, 3, 2, 8, 4, 9},
			fee:    2,
			result: 8,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, maxProfit(tc.nums, tc.fee))
	}
}
