package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		nums   int
		result int
	}{
		{
			nums:   4421,
			result: 21,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, subtractProductAndSum(tc.nums))
	}
}
