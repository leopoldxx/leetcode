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
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			result: []int{3, 4},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, searchRange(tc.nums, tc.target))
	}
}
