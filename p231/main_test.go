package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums1  []int
		nums2  []int
		result int
	}{
		{
			nums1:  []int{2, 4},
			result: 3,
		},
		{
			nums1:  []int{2, 4, 5, 10},
			result: 7,
		},
		{
			nums1:  []int{18, 3, 6, 2},
			result: 12,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, numFactoredBinaryTrees(tc.nums1))
	}
}
