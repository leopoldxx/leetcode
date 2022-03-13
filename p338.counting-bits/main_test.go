package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		num    int
		result []int
	}{
		{
			num:    2,
			result: []int{0, 1, 1},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, countBits(tc.num))
	}
}
