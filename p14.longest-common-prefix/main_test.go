package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		strs   []string
		result string
	}{
		{
			strs:   []string{"flower", "flow", "flight"},
			result: "fl",
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, longestCommonPrefix(tc.strs))
	}
}
