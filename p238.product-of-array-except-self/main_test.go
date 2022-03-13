package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums1  []int
		result []int
	}{
		{
			nums1:  []int{1, 2, 3, 4},
			result: []int{24, 12, 8, 6},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, productExceptSelf(tc.nums1))
	}
}
