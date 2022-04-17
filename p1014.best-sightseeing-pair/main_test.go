package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		m      []int
		result int
	}{
		{
			m:      []int{8, 1, 5, 2, 6},
			result: 11,
		},
		{
			m:      []int{1, 2},
			result: 2,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, maxScoreSightseeingPair(tc.m))
	}
}
