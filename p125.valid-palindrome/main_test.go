package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		s      string
		result bool
	}{
		{
			s:      "A man, a plan, a canal: Panama",
			result: true,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, isPalindrome(tc.s))
	}
}
