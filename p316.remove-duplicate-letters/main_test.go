package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		s      string
		result string
	}{
		{
			s:      "cbacdcbc",
			result: "acdb",
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, removeDuplicateLetters(tc.s))
	}
}
