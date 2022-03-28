package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		grid   [][]byte
		result int
	}{
		{
			grid:   [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}},
			result: 2,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, numIslands(tc.grid))
	}
}
