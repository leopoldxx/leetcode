package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums   []int
		val    int
		result int
	}{
		{
			nums:   []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:    2,
			result: 5,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, removeElement(tc.nums, tc.val))
	}
}
