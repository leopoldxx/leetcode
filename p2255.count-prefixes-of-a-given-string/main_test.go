package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 2, countPrefixes([]string{"a", "ab", "c"}, "abc"))
}
