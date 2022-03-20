package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		intervals [][]int
		result    int
	}{
		{
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			result:    1,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, eraseOverlapIntervals(tc.intervals))
	}
}
