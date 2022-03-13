package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		num    int
		result bool
	}{
		{
			num:    2,
			result: false,
		},
		{
			num:    3,
			result: true,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, isPowerOfThree(tc.num))
	}
}
