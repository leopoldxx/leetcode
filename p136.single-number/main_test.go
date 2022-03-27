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
			nums:   []int{2, 2, 1},
			result: 1,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, singleNumber(tc.nums))
	}
}
