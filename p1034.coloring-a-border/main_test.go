package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		grid   [][]int
		row    int
		col    int
		color  int
		result [][]int
	}{
		{
			grid:   [][]int{{1, 1}, {1, 2}},
			row:    0,
			col:    0,
			color:  3,
			result: [][]int{{3, 3}, {3, 2}},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, colorBorder(tc.grid, tc.row, tc.col, tc.color))
	}
}
