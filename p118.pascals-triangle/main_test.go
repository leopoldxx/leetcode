package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		numRows int
		result  [][]int
	}{
		{
			numRows: 5,
			result: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, generate(tc.numRows))
	}
}
