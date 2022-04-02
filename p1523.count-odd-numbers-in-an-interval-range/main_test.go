package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		low    int
		high   int
		result int
	}{
		{
			low:    3,
			high:   7,
			result: 3,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, countOdds(tc.low, tc.high))
	}
}
