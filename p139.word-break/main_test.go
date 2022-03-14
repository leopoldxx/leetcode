package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		s        string
		wordDict []string
		result   bool
	}{

		{
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			result:   true,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, wordBreak(tc.s, tc.wordDict))
	}
}
