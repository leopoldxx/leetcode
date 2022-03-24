package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		s      string
		k      int
		result int
	}{

		{
			s:      "BAA",
			k:      0,
			result: 2,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, characterReplacement(tc.s, tc.k))
	}
}
