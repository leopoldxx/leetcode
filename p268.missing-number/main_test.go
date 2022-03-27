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
			nums:   []int{3, 0, 1},
			result: 2,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, missingNumber(tc.nums))
	}
}
