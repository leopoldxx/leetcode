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
			nums:   []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			result: 5,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, removeDuplicates(tc.nums))
	}
}
