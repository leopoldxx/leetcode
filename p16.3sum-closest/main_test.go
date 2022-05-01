package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 3, threeSumClosest([]int{0, 1, 2}, 3))
}
