package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.ElementsMatch(t, []string{"a1b2", "a1B2", "A1b2", "A1B2"}, letterCasePermutation(string("a1b2")))
}
