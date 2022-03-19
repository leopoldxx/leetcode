package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		m      int
		n      int
		result int
	}{
		{
			m:      3,
			n:      7,
			result: 28,
		},
		{
			m:      3,
			n:      2,
			result: 3,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, uniquePaths(tc.m, tc.n))
	}
}
