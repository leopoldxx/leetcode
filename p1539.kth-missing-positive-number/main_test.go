package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums   []int
		k      int
		result int
	}{
		{
			nums:   []int{2, 3, 4, 7, 11},
			k:      5,
			result: 9,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, findKthPositive(tc.nums, tc.k))
	}
}
