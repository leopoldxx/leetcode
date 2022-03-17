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
			s:      "(()(()))",
			result: 6,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, scoreOfParentheses(tc.s))
	}
}
