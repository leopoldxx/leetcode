package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, true, searchMatrix([][]int{{1}}, 1))
}
