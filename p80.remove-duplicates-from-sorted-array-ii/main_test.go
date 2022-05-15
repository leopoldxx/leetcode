package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 5, removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
}
