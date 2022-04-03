package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		result []int
	}{
		{
			nums:   []int{1, 2, 5, 2, 3},
			target: 3,
			result: []int{3},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, targetIndices(tc.nums, tc.target))
	}
}
