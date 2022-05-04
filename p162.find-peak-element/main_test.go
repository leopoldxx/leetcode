package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 2, findPeakElement([]int{1, 2, 3, 1}))
}
