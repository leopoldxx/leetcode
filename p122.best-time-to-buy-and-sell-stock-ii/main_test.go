package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		m      []int
		result int
	}{
		{
			m:      []int{7, 1, 5, 3, 6, 4},
			result: 7,
		},
		{
			m:      []int{1, 2, 3, 4, 5},
			result: 4,
		},
		{
			m:      []int{7, 6, 4, 3, 1},
			result: 0,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, maxProfit(tc.m))
	}
}
