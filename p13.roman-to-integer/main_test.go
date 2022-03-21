package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		s      string
		result int
	}{
		{
			s:      "MCMXCIV",
			result: 1994,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, romanToInt(tc.s))
	}
}
