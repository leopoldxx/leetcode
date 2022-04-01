package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		s      []int
		result []int
	}{
		{
			s:      []int{1, 2, 3},
			result: []int{1, 2, 4},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, plusOne(tc.s))
	}
}
