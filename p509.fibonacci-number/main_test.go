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
			num:    2,
			result: 1,
		},
		{
			num:    4,
			result: 3,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, fib(tc.num))
	}
}
