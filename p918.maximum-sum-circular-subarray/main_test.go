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
			nums:   []int{1, -2, 3, -2},
			result: 3,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, maxSubarraySumCircular(tc.nums))
	}
}
