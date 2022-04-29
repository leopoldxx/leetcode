package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, letterCombinations("2"))
}
