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
			nums:   []int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6},
			result: 23,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, trap(tc.nums))
	}
}
