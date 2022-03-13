package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums1  []int
		result bool
	}{
		{
			nums1:  []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			result: true,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, containsDuplicate(tc.nums1))
	}
}
