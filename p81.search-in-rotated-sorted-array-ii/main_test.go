package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, true, search([]int{1, 0, 1, 1, 1}, 0))
	assert.Equal(t, true, search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
}
