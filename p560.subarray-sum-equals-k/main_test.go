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
			nums:   []int{1, 1, 1},
			target: 2,
			result: 2,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, subarraySum(tc.nums, tc.target))
	}
}
