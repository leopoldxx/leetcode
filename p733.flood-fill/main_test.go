package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		image  [][]int
		x      int
		y      int
		color  int
		result [][]int
	}{
		{
			image:  [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			x:      1,
			y:      1,
			color:  2,
			result: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, floodFill(tc.image, tc.x, tc.y, tc.color))
	}
}
