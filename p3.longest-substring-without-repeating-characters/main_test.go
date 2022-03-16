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
			s:      "abcabcbb",
			result: 3,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, lengthOfLongestSubstring(tc.s))
	}
}
