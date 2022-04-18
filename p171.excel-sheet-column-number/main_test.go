package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		columnTitle string
		result      int
	}{
		{
			columnTitle: "AB",
			result:      28,
		},
		{
			columnTitle: "A",
			result:      1,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, titleToNumber(tc.columnTitle))
	}
}
