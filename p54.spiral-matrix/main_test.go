package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		matrix [][]int
		result []int
	}{
		{
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			result: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			result: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, spiralOrder(tc.matrix))
	}

}
