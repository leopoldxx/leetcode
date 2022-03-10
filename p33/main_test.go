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
			nums:   []int{5, 1, 3},
			target: 5,
			result: 0,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, search(tc.nums, tc.target))
	}
}
