package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 17, maximumUniqueSubarray([]int{4, 2, 4, 5, 6}))
}
