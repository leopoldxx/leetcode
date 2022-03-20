package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		intervals [][]int
		result    [][]int
	}{
		{
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			result:    [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, merge(tc.intervals))
	}
}
