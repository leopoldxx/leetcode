package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums1  []int
		m      int
		nums2  []int
		n      int
		result []int
	}{
		{
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{2, 5, 6},
			n:      3,
			result: []int{1, 2, 2, 3, 5, 6},
		},
	}
	for _, tc := range cases {
		merge(tc.nums1, tc.m, tc.nums2, tc.n)
		assert.Equal(t, tc.result, tc.nums1)
	}
}
