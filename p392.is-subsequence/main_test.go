package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.True(t, isSubsequence("ana", "acnda"))
	assert.False(t, isSubsequence("anc", "acnda"))
}
