package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums   uint32
		result int
	}{
		{
			nums:   3,
			result: 2,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, hammingWeight(tc.nums))
	}
}
