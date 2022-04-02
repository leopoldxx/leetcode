package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		salary []int
		result float64
	}{
		{
			salary: []int{4000, 3000, 1000, 2000},
			result: 2500.0,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.result, average(tc.salary))
	}
}
