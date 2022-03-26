package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		result int
	}{
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			result: 4,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			result: -1,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, search(tc.nums, tc.target))
	}
}
