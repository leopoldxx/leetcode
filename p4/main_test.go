package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	cases := []struct {
		nums1  []int
		nums2  []int
		result float64
	}{
		{
			nums1:  []int{1, 3},
			nums2:  []int{2},
			result: 2,
		},
		{
			nums1:  []int{1, 2},
			nums2:  []int{3, 4},
			result: 2.5,
		},
		{

			nums1:  []int{2, 2, 4, 4},
			nums2:  []int{2, 2, 4, 4},
			result: 3,
		},
		{

			nums1:  []int{1},
			nums2:  []int{2, 3, 4, 5, 6, 7},
			result: 4,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, findMedianSortedArrays(tc.nums1, tc.nums2))
	}
}
