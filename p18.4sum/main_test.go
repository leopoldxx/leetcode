package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, [][]int{{2, 2, 2, 2}}, fourSum([]int{2, 2, 2, 2, 2}, 8))
}
