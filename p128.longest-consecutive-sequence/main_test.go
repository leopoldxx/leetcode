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
			nums:   []int{100, 4, 200, 1, 3, 2},
			result: 4,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, longestConsecutive(tc.nums))
	}
}
