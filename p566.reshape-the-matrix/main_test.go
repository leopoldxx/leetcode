package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		mat    [][]int
		r      int
		c      int
		result [][]int
	}{
		{
			mat:    [][]int{{1, 2}, {3, 4}},
			r:      1,
			c:      4,
			result: [][]int{{1, 2, 3, 4}},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, matrixReshape(tc.mat, tc.r, tc.c))
	}
}
