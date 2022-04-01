package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		a      string
		b      string
		result string
	}{
		{
			a:      "11",
			b:      "1",
			result: "100",
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, addBinary(tc.a, tc.b))
	}
}
