package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums   []int
		k      int
		result []int
	}{
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			k:      3,
			result: []int{5, 6, 7, 1, 2, 3, 4},
		},
	}
	for _, tc := range cases {
		rotate(tc.nums, tc.k)
		assert.Equal(t, tc.result, tc.nums)
	}
}
