package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		grid   [][]int
		result int
	}{
		{
			grid:   [][]int{{1}},
			result: 1,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, maxAreaOfIsland(tc.grid))
	}
}
