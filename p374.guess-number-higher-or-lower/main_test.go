package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	guessFunc := func(result int) func(num int) int {
		return func(num int) int {
			if num == result {
				return 0
			} else if num > result {
				return -1
			} else {
				return 1
			}
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
		guess = guessFunc(tc.result)
		assert.Equal(t, tc.result, guessNumber(tc.num))
	}
}
