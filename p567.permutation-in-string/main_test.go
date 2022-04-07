package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		s1     string
		s2     string
		result bool
	}{
		{
			s1:     "ab",
			s2:     "bab",
			result: true,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, checkInclusion(tc.s1, tc.s2))
	}
}
