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
			nums:   []int{2, 3, 1, 1, 4},
			result: true,
		},
		{
			nums:   []int{3, 2, 1, 0, 4},
			result: false,
		},
		{
			nums:   []int{0},
			result: true,
		},
		{
			nums:   []int{0, 1},
			result: false,
		},
		{
			nums:   []int{2, 0, 0},
			result: true,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, canJump(tc.nums))
	}
}
