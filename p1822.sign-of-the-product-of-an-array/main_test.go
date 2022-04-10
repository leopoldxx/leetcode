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
			nums:   []int{2, 3, 4, 7, 11},
			result: 1,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, arraySign(tc.nums))
	}
}
