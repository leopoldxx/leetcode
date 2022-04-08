package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		x      int
		y      int
		points [][]int
		result int
	}{
		{
			x:      3,
			y:      4,
			points: [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}},
			result: 2,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, nearestValidPoint(tc.x, tc.y, tc.points))
	}
}
