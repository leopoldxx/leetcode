package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		x      int
		result int
	}{
		{
			x:      4,
			result: 2,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, mySqrt(tc.x))
	}
}
