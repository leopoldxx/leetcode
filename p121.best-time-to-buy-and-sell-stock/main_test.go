package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums1  []int
		result int
	}{
		{
			nums1:  []int{7, 1, 5, 3, 6, 4},
			result: 5,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, maxProfit(tc.nums1))
	}
}
