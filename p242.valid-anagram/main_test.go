package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		s      string
		t      string
		result bool
	}{
		{
			s:      "anagram",
			t:      "nagaram",
			result: true,
		},
		{
			s:      "cat",
			t:      "car",
			result: false,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, isAnagram(tc.s, tc.t))
	}

}
