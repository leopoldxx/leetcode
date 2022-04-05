package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		letters []byte
		target  byte
		result  byte
	}{
		{
			letters: []byte{'e', 'e', 'e', 'e', 'e', 'e', 'n', 'n', 'n', 'n'},
			target:  'e',
			result:  'n',
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, nextGreatestLetter(tc.letters, tc.target))
	}
}
