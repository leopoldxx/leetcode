package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		text1  string
		text2  string
		result int
	}{

		{
			text1:  "bsbininm",
			text2:  "jmjkbkjkv",
			result: 1,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, longestCommonSubsequence(tc.text1, tc.text2))
	}
}
