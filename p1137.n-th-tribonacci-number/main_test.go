package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		num    int
		result int
	}{
		{
			num:    4,
			result: 4,
		},
		{
			num:    25,
			result: 1389537,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, tribonacci(tc.num))
	}
}
