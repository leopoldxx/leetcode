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
			nums:   []int{1, 2, 3},
			target: 4,
			result: 7,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, combinationSum4(tc.nums, tc.target))
	}
}
