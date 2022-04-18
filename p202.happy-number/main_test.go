package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		num    int
		result bool
	}{
		{
			num:    1111111,
			result: true,
		},
		{
			num:    2,
			result: false,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, isHappy(tc.num))
	}
}
