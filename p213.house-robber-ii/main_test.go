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
			nums:   []int{200, 3, 140, 20, 10},
			result: 340,
		},
		{
			nums:   []int{1, 2, 3, 1},
			result: 4,
		},
		{
			nums:   []int{1, 2, 3},
			result: 3,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, rob(tc.nums))
	}
}
