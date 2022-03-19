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
			s:      "226",
			result: 3,
		},
		{
			s:      "12",
			result: 2,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, numDecodings(tc.s))
	}
}
