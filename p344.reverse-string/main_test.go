package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		s      []byte
		result []byte
	}{
		{
			s:      []byte{'h', 'e', 'l', 'l', 'o'},
			result: []byte{'o', 'l', 'l', 'e', 'h'},
		},
	}
	for _, tc := range cases {
		reverseString(tc.s)
		assert.Equal(t, tc.result, tc.s)
	}
}
