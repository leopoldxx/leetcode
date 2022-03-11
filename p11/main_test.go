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
			nums:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			result: 49,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, maxArea(tc.nums))
	}
}
