package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.ElementsMatch(t, [][]int{
		{1, 1, 2},
		{1, 2, 1},
		{2, 1, 1},
	}, permuteUnique([]int{1, 1, 2}))
}
