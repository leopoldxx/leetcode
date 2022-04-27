package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, "dqwyt", smallestStringWithSwaps("qdwyt", [][]int{{2, 3}, {3, 2}, {0, 1}, {4, 0}, {3, 2}}))
	assert.Equal(t, "bacd", smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}}))
}
