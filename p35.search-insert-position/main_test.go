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
			nums:   []int{1, 3, 5, 6},
			target: 5,
			result: 2,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, searchInsert(tc.nums, tc.target))
	}
}
