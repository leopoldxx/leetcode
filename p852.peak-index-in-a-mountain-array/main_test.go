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
			nums:   []int{0, 10, 5, 2},
			result: 1,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, peakIndexInMountainArray(tc.nums))
	}
}
