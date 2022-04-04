package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		arr1   []int
		arr2   []int
		target int
		result int
	}{
		{
			arr1:   []int{-3, 10, 2, 8, 0, 10},
			arr2:   []int{-9, -1, -4, -9, -8},
			target: 9,
			result: 2,
		},
		{
			arr1:   []int{1, 4, 2, 3},
			arr2:   []int{-4, -3, 6, 10, 20, 30},
			target: 3,
			result: 2,
		},
		{
			arr1:   []int{4, 5, 8},
			arr2:   []int{10, 9, 1, 8},
			target: 2,
			result: 2,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, findTheDistanceValue(tc.arr1, tc.arr2, tc.target))
	}
}
