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
			nums:   []int{2, 7, 9, 3, 1},
			result: 12,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, rob(tc.nums))
	}
}
