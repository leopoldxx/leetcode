package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		matrix [][]int
		result [][]int
	}{
		{
			matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			result: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}
	for _, tc := range cases {
		setZeroes(tc.matrix)
		assert.Equal(t, tc.result, tc.matrix)
	}
}
