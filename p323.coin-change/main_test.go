package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		coins  []int
		amount int
		result int
	}{
		{
			coins:  []int{1, 2, 5},
			amount: 11,
			result: 3,
		},
		{
			coins:  []int{2},
			amount: 3,
			result: -1,
		},
		{
			coins:  []int{1},
			amount: 0,
			result: 0,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, coinChange(tc.coins, tc.amount))
	}
}
