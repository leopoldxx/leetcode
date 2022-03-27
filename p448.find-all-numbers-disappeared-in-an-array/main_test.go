package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums   []int
		result []int
	}{
		{
			nums:   []int{2, 2},
			result: []int{1},
		},
		{
			nums:   []int{4, 3, 2, 7, 8, 2, 3, 1},
			result: []int{5, 6},
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, findDisappearedNumbers(tc.nums))
	}
}
