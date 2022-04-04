package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums1  []int
		nums2  []int
		result []int
	}{
		{
			nums1:  []int{1, 2, 2, 1},
			nums2:  []int{2, 2},
			result: []int{2, 2},
		},
		{
			nums1:  []int{4, 9, 5},
			nums2:  []int{9, 4, 9, 8, 4},
			result: []int{9, 4},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, intersect(tc.nums1, tc.nums2))
	}
}
