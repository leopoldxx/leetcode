package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums   []int
		result bool
	}{
		{
			nums:   []int{0, 5, 3, 4, 6},
			result: false,
		},
		{
			nums:   []int{3, 5, 1},
			result: true,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, canMakeArithmeticProgression(tc.nums))
	}
}
