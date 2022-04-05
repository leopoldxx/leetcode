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
			s:      "Let's take LeetCode contest",
			result: "s'teL ekat edoCteeL tsetnoc",
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, reverseWords(tc.s))
	}
}
