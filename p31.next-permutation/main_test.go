package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums   []int
		result []int
	}{
		{
			nums:   []int{1, 2, 3},
			result: []int{1, 3, 2},
		},
		{
			nums:   []int{1, 3, 2},
			result: []int{2, 1, 3},
		},
	}
	for _, tc := range cases {
		nextPermutation(tc.nums)
		assert.Equal(t, tc.result, tc.nums)
	}
}
