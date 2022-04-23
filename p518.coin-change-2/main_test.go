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
			amount: 5,
			result: 4,
		},
		{
			coins:  []int{2},
			amount: 3,
			result: 0,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, change(tc.amount, tc.coins))
	}
}
