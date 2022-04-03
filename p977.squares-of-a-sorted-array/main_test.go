package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		salary []int
		result []int
	}{
		{
			salary: []int{-4, -1, 0, 3, 10},
			result: []int{0, 1, 9, 16, 100},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, sortedSquares(tc.salary))
	}
}
