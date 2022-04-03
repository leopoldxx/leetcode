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
			nums:   []int{0, 10, 5, 2},
			k:      1,
			result: 10,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, findKthLargest(tc.nums, tc.k))
	}
}
