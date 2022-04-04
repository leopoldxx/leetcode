package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		x      int
		result bool
	}{
		{
			x:      16,
			result: true,
		},
		{
			x:      14,
			result: false,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, isPerfectSquare(tc.x))
	}
}
