package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		num    uint32
		result uint32
	}{
		{
			num:    3,
			result: 0xc0000000,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, reverseBits(tc.num))
	}
}
