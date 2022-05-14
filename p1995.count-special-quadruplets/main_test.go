package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 1, countQuadruplets([]int{28, 8, 49, 85, 37, 90, 20, 8}))
}
