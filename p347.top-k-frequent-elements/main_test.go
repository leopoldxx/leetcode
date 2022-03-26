package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums   []int
		k      int
		result []int
	}{
		{
			nums:   []int{1, 1, 1, 2, 2, 3},
			k:      2,
			result: []int{1, 2},
		},
	}
	for _, tc := range cases {
		assert.ElementsMatch(t, tc.result, topKFrequent(tc.nums, tc.k))
	}
}
