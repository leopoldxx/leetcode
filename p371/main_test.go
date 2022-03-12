package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		a, b   int
		result int
	}{
		{
			a:      -4,
			b:      1,
			result: -3,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, getSum(tc.a, tc.b))
	}
}
