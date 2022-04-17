package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		mat    [][]int
		k      int
		result [][]int
	}{
		{
			mat:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:      1,
			result: [][]int{{12, 21, 16}, {27, 45, 33}, {24, 39, 28}},
		},
		{
			mat:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:      2,
			result: [][]int{{45, 45, 45}, {45, 45, 45}, {45, 45, 45}},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, matrixBlockSum(tc.mat, tc.k))
	}
}
