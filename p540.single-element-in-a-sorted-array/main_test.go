package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums   []int
		result int
	}{
		{
			nums:   []int{3, 3, 7, 7, 10, 11, 11},
			result: 10,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, singleNonDuplicate(tc.nums))
	}
}
