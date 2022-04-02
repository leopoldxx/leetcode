package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	checkerFunc := func(result int) func(num int) bool {
		return func(num int) bool {
			return num == result
		}
	}
	cases := []struct {
		num    int
		result int
	}{
		{
			num:    5,
			result: 4,
		},
		{
			num:    1,
			result: 1,
		},
	}
	for _, tc := range cases {
		isBadVersion = checkerFunc(tc.result)
		assert.Equal(t, tc.result, firstBadVersion(tc.num))
	}
}
