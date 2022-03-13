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
			nums:   []int{10, 9, 2, 5, 3, 7, 101, 18},
			result: 4,
		},
		{
			nums:   []int{7, 7, 7, 7, 7, 7, 7},
			result: 1,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, lengthOfLIS(tc.nums))
	}
}
