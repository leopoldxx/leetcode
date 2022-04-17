package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		m      int
		k      int
		result string
	}{
		{
			m:      3,
			k:      2,
			result: "132",
		},
		{
			m:      3,
			k:      3,
			result: "213",
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, getPermutation(tc.m, tc.k))
	}
}
