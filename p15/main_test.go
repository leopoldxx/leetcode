package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums1  []int
		result [][]int
	}{
		{
			nums1:  []int{-1, 0, 1, 2, -1, -4},
			result: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, threeSum(tc.nums1))
	}
}
