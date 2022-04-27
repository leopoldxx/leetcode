package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		numRows int
		result  []int
	}{
		{
			numRows: 4,
			result:  []int{1, 4, 6, 4, 1},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, getRow(tc.numRows))
	}
}
