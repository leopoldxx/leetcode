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
			nums:   []int{2, 3, 1, 1, 4},
			result: 2,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, jump(tc.nums))
	}
}
