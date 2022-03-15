package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		intervals   [][]int
		newInterval []int
		result      [][]int
	}{
		{
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			result:      [][]int{{1, 5}, {6, 9}},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, insert(tc.intervals, tc.newInterval))
	}
}
